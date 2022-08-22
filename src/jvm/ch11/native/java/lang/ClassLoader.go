package lang

import (
  "go-jvm/src/jvm/ch11/native"
  "go-jvm/src/jvm/ch11/rtda"
)

func init() {
  native.Register("java/lang/ClassLoader", "findBuiltinLib", "(Ljava/lang/String;)Ljava/lang/String;", findBuiltinLib)
}

func findBuiltinLib(frame *rtda.Frame) {
  frame.OperandStack().PushRef(nil)
}
