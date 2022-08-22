package constants

import (
	"go-jvm/src/jvm/ch09/instructions/base"
	"go-jvm/src/jvm/ch09/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// to do nothing
}
