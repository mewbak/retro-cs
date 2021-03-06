package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/blackchip-org/retro-cs/app/monitor"
	"github.com/blackchip-org/retro-cs/mock"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/blackchip-org/retro-cs/app"
	"github.com/blackchip-org/retro-cs/config"
	"github.com/blackchip-org/retro-cs/rcs"
)

const (
	defaultWidth  = 1024
	defaultHeight = 786
)

var (
	optFullStart bool
	optProfC     bool
	optPanic     bool
	optSystem    string
	optMonitor   bool
	optImport    string
	optNoAudio   bool
	optNoVideo   bool
	optTrace     bool
	optWait      bool
)

func init() {
	flag.BoolVar(&optFullStart, "f", false, "full start -- do not bypass POST")
	flag.StringVar(&optImport, "i", "", "import state from `filename`")
	flag.BoolVar(&optProfC, "profc", false, "enable cpu profiling")
	flag.BoolVar(&optNoAudio, "no-audio", false, "disable audio")
	flag.BoolVar(&optNoVideo, "no-video", false, "disable video")
	flag.BoolVar(&optMonitor, "m", false, "enable monitor")
	flag.BoolVar(&optPanic, "panic", false, "install panic log writer")
	flag.StringVar(&optSystem, "s", "c64", "start this `system`")
	flag.BoolVar(&optTrace, "t", false, "enable tracing")
	flag.BoolVar(&optWait, "w", false, "wait for go command")
}

func main() {
	runtime.LockOSThread()
	log.SetFlags(0)
	flag.Parse()

	if optProfC {
		f, err := os.Create("./cpu.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		log.Println("starting profile")
		defer func() {
			pprof.StopCPUProfile()
			log.Println("profile saved")
		}()
	}
	if optNoVideo || optTrace || optWait {
		optMonitor = true
	}

	newMachine, ok := app.Systems[optSystem]
	if !ok {
		log.Fatalf("no such system: %v", optSystem)
	}
	config.System = optSystem
	config.UserDir = filepath.Join(config.UserHome, ".retro-cs")
	config.DataDir = filepath.Join(config.ResourceDir(), "data", optSystem)
	config.VarDir = filepath.Join(config.ResourceDir(), "var", optSystem)

	if err := os.MkdirAll(config.UserDir, 0755); err != nil {
		log.Fatalf("unable to create directory %v: %v", config.UserDir, err)
	}
	if err := os.MkdirAll(config.VarDir, 0755); err != nil {
		log.Fatalf("unable to create directory %v: %v", config.VarDir, err)
	}

	ctx := rcs.SDLContext{}
	if !optNoVideo {
		if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
			log.Fatalf("unable to initialize video: %v", err)
		}
		fullScreen := uint32(0)
		if !optMonitor {
			fullScreen = sdl.WINDOW_FULLSCREEN_DESKTOP
		}
		window, err := sdl.CreateWindow(
			"retro-cs",
			sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			defaultWidth, defaultHeight,
			sdl.WINDOW_SHOWN|fullScreen,
		)
		if err != nil {
			log.Fatalf("unable to initialize window: %v", err)
		}
		r, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			log.Fatalf("unable to initialize renderer: %v", err)
		}
		info, err := r.GetInfo()
		if err != nil {
			log.Fatalf("unable to get renderer info: %v", err)
		}
		// FIXME: Change this to check for OpenGL
		if info.Name != "direct3d" {
			if _, err := window.GLCreateContext(); err != nil {
				log.Printf("unable to createl GL context: %v", err)
			}
			if err = sdl.GLSetSwapInterval(1); err != nil {
				log.Printf("unable to set swap interval: %v", err)
			}
		}
		ctx.Window = window
		ctx.Renderer = r
	}

	if !optNoAudio {
		requestSpec := sdl.AudioSpec{
			Freq:     22050,
			Format:   sdl.AUDIO_S16LSB,
			Channels: 2,
			Samples:  367,
		}
		if err := sdl.OpenAudio(&requestSpec, &ctx.AudioSpec); err != nil {
			log.Fatalf("unable to initialize audio: %v", err)
		}
		sdl.PauseAudio(false)
	}

	err := sdl.Init(sdl.INIT_JOYSTICK | sdl.INIT_GAMECONTROLLER)
	if err != nil {
		log.Fatalf("unable to initialize game controllers: %v", err)
	}
	mappingsDir := filepath.Join(config.ResourceDir(), "data", "game_controllers")
	mappings, err := ioutil.ReadDir(mappingsDir)
	if err != nil {
		log.Printf("(!) unable to load game controller mappings: %v", err)
	}
	for _, f := range mappings {
		if strings.HasSuffix(f.Name(), ".txt") {
			in, err := os.Open(filepath.Join(mappingsDir, f.Name()))
			if err != nil {
				log.Printf("(!) unable to open controller mapping: %v", err)
				continue
			}
			defer in.Close()
			s := bufio.NewScanner(in)
			for s.Scan() {
				line := strings.TrimSpace(s.Text())
				if line == "" {
					continue
				}
				if line[0] == '#' {
					continue
				}
				sdl.GameControllerAddMapping(line)
			}
		}
	}

	mach, err := newMachine(ctx)
	if err != nil {
		log.Fatalf("unable to create machine: \n%v", err)
	}

	var mon *monitor.Monitor
	mon, err = monitor.New(mach)
	if err != nil {
		log.Fatalf("unable to create monitor: %v\n", err)
	}
	defer func() {
		mon.Close()
	}()

	if optMonitor {
		go func() {
			err := mon.Run()
			if err != nil {
				log.Fatalf("monitor error: %v", err)
			}
		}()
	}

	mach.Status = rcs.Run
	if optWait {
		mach.Status = rcs.Pause
	}
	if optTrace {
		mach.Command(rcs.MachTraceAll, true)
	}
	if optImport != "" {
		filename := filepath.Join(config.VarDir, optImport)
		mach.Command(rcs.MachImport, filename)
	} else if !optFullStart {
		filename := filepath.Join(config.DataDir, "init.state")
		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			mach.Command(rcs.MachImport, filename)
		}
	}

	if optPanic {
		log.SetOutput(&mock.PanicWriter{})
	}

	startFile := filepath.Join(config.UserDir, "startup")
	cmds, err := ioutil.ReadFile(startFile)
	if err == nil {
		mon.Eval(string(cmds))
	}

	mach.Run()
}
