package constants

import (
	"go-jvm/src/jvm/ch10/instructions/base"
	"go-jvm/src/jvm/ch10/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// to do nothing
}
