package lang

import (
  "go-jvm/src/jvm/ch09/native"
  "go-jvm/src/jvm/ch09/rtda"
  "unsafe"
)

const jlObject = "java/lang/Object"

func init() {
  native.Register(jlObject, "getClass", "()Ljava/lang/Class;", getClass)
  native.Register(jlObject, "hashCode", "()I", hashCode)
  native.Register(jlObject, "clone", "()Ljava/lang/Object;", clone)
}

// public final native Class<?> getClass();
func getClass(frame *rtda.Frame) {
  this := frame.LocalVars().GetThis()
  class := this.Class().JClass()
  frame.OperandStack().PushRef(class)
}

// public native int hashCode();
func hashCode(frame *rtda.Frame) {
  this := frame.LocalVars().GetThis()
  hash := int32(uintptr(unsafe.Pointer(this)))
  frame.OperandStack().PushInt(hash)
}

func clone(frame *rtda.Frame) {
  this := frame.LocalVars().GetThis()
  cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
  if !this.Class().IsImplements(cloneable) {
    panic("java.lang.CloneNotSupportedException")
  }
  frame.OperandStack().PushRef(this.Clone())
}
