package references

import (
  "go-jvm/src/jvm/ch07/instructions/base"
  "go-jvm/src/jvm/ch07/rtda"
  "go-jvm/src/jvm/ch07/rtda/heap"
)

// Create new object
type NEW struct{ base.Index16Instruction }

func (self *NEW) Execute(frame *rtda.Frame) {
  cp := frame.Method().Class().ConstantPool()
  classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
  class := classRef.ResolvedClass()
  // todo: init class
  if !class.InitStarted() {
    frame.RevertNextPC()
    base.InitClass(frame.Thread(), class)
    return
  }

  if class.IsInterface() || class.IsAbstract() {
    panic("java.lang.InstantiationError")
  }

  ref := class.NewObject()
  frame.OperandStack().PushRef(ref)
}
