package control

import (
	"go-jvm/src/jvm/ch05/instructions/base"
	"go-jvm/src/jvm/ch05/rtda"
)

// Branch always
type GOTO struct{ base.BranchInstruction }



func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}
