package constants

import (
	"go-jvm/src/jvm/ch05/instructions/base"
	"go-jvm/src/jvm/ch05/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// to do nothing
}
