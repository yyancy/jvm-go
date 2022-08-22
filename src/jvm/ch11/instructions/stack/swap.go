package stack

import (
	"go-jvm/src/jvm/ch11/instructions/base"
	"go-jvm/src/jvm/ch11/rtda"
)

type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
