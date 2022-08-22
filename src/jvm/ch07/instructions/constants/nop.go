package constants

import (
	"go-jvm/src/jvm/ch07/instructions/base"
	"go-jvm/src/jvm/ch07/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// to do nothing
}
