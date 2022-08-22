package lang

import (
  "go-jvm/src/jvm/ch11/native"
  "go-jvm/src/jvm/ch11/rtda"
  "runtime"
)

const jlRuntime = "java/lang/Runtime"

func init() {
  native.Register(jlRuntime, "availableProcessors", "()I", availableProcessors)
}

// public native int availableProcessors();
// ()I
func availableProcessors(frame *rtda.Frame) {
  numCPU := runtime.NumCPU()

  stack := frame.OperandStack()
  stack.PushInt(int32(numCPU))
}
