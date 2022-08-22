package references

import (
	"go-jvm/src/jvm/ch06/instructions/base"
	"go-jvm/src/jvm/ch06/rtda"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
