package extended

import (
	"go-jvm/src/jvm/ch05/instructions/base"
	"go-jvm/src/jvm/ch05/rtda"
)

// Branch if reference is null
type IFNULL struct{ base.BranchInstruction }

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ base.BranchInstruction }

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
