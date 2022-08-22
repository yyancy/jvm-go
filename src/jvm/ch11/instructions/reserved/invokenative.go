package reserved

import (
  "go-jvm/src/jvm/ch11/instructions/base"
  "go-jvm/src/jvm/ch11/native"
  _ "go-jvm/src/jvm/ch11/native/java/io"
  _ "go-jvm/src/jvm/ch11/native/java/lang"
  _ "go-jvm/src/jvm/ch11/native/java/security"
  _ "go-jvm/src/jvm/ch11/native/java/util/concurrent/atomic"
  _ "go-jvm/src/jvm/ch11/native/sun/io"
  _ "go-jvm/src/jvm/ch11/native/sun/misc"
  _ "go-jvm/src/jvm/ch11/native/sun/reflect"
  "go-jvm/src/jvm/ch11/rtda"
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
