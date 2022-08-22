package reserved

import (
  "go-jvm/src/jvm/ch09/instructions/base"
  "go-jvm/src/jvm/ch09/native"
  _ "go-jvm/src/jvm/ch09/native/java/lang"
  _ "go-jvm/src/jvm/ch09/native/sun/misc"
  "go-jvm/src/jvm/ch09/rtda"
)

type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
  method := frame.Method()
  className := method.Class().Name()
  methodName := method.Name()
  methodDescriptor := method.Descriptor()

  nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
  if nativeMethod == nil {
    methodInfo := className + "." + methodName + methodDescriptor
    panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
  }

  nativeMethod(frame)
}
