package m6502

import (
	"github.com/blackchip-org/retro-cs/rcs"
)

// add with carry
func adc(c *CPU, load rcs.Load8) {
	if c.SR&FlagD != 0 {
		c.A = c.alu.AddBCD(&c.SR, c.A, load())
	} else {
		c.A = c.alu.Add(&c.SR, c.A, load())
	}
}

// logical and
func and(c *CPU, load rcs.Load8) {
	c.A = c.alu.And(&c.SR, c.A, load())
}

// arithmetic shift left
func asl(c *CPU, store rcs.Store8, load rcs.Load8) {
	c.SR &^= FlagC // shift in zero
	store(c.alu.ShiftLeft(&c.SR, load()))
}

// test bits
func bit(c *CPU, load rcs.Load8) {
	in := load()

	if c.A&in == 0 {
		c.SR |= FlagZ
	} else {
		c.SR &^= FlagZ
	}

	if in&(1<<7) != 0 {
		c.SR |= FlagN
	} else {
		c.SR &^= FlagN
	}

	if in&(1<<6) != 0 {
		c.SR |= FlagV
	} else {
		c.SR &^= FlagV
	}
}

// branch instructions
func branch(c *CPU, do bool) {
	displacement := int8(c.fetch())
	if do {
		if displacement >= 0 {
			c.SetPC(c.PC() + int(displacement))
		} else {
			c.SetPC(c.PC() - int(displacement*-1))
		}
	}
}

// break
func brk(c *CPU) {
	c.SR |= FlagB
	c.fetch()
}

// compare
func cmp(c *CPU, load0 rcs.Load8, load1 rcs.Load8) {
	// C set as if subtraction. Clear if 'borrow', otherwise set
	r := int16(load0()) - int16(load1())
	if r >= 0 {
		c.SR |= FlagC
	} else {
		c.SR &^= FlagC
	}
	c.alu.Pass(&c.SR, uint8(r))
}

// decrement
func dec(c *CPU, store rcs.Store8, load rcs.Load8) {
	store(c.alu.Decrement(&c.SR, load()))
}

// exclusive or
func eor(c *CPU, load rcs.Load8) {
	c.A = c.alu.ExclusiveOr(&c.SR, c.A, load())
}

// increment
func inc(c *CPU, store rcs.Store8, load rcs.Load8) {
	store(c.alu.Increment(&c.SR, load()))
}

// jump
func jmp(c *CPU) {
	c.pc = uint16(c.fetch2() - 1)
}

// jump indirect
func jmpIndirect(c *CPU) {
	c.pc = uint16(c.mem.ReadLE(c.fetch2()) - 1)
}

// jump to subroutine
func jsr(c *CPU) {
	addr := uint16(c.fetch2())
	c.push2(c.pc)
	c.pc = addr - 1
}

// load
func ld(c *CPU, store rcs.Store8, load rcs.Load8) {
	value := load()
	store(value)
	c.alu.Pass(&c.SR, value)
}

// logical shift right
func lsr(c *CPU, store rcs.Store8, load rcs.Load8) {
	c.SR &^= FlagC // shift in zero
	store(c.alu.ShiftRight(&c.SR, load()))
}

// logical or
func ora(c *CPU, load rcs.Load8) {
	c.A = c.alu.Or(&c.SR, c.A, load())
}

// push processor status
func php(c *CPU) {
	// https://wiki.nesdev.com/w/index.php/Status_flags
	c.push(c.SR | FlagB)
}

// pull accumulator
func pla(c *CPU) {
	c.A = c.pull()
	c.alu.Pass(&c.SR, c.A)
}

// rotate left
func rol(c *CPU, store rcs.Store8, load rcs.Load8) {
	store(c.alu.ShiftLeft(&c.SR, load()))
}

// rotate right
func ror(c *CPU, store rcs.Store8, load rcs.Load8) {
	store(c.alu.ShiftRight(&c.SR, load()))
}

// return from interrupt
func rti(c *CPU) {
	c.SR = c.pull()
	c.pc = c.pull2() - 1
}

// subtract with carry
func sbc(c *CPU, load rcs.Load8) {
	if c.SR&FlagD != 0 {
		c.A = c.alu.SubtractBCD(&c.SR, c.A, load())
	} else {
		c.A = c.alu.Subtract(&c.SR, c.A, load())
	}
}

// store
func st(c *CPU, store rcs.Store8, load rcs.Load8) {
	store(load())
}