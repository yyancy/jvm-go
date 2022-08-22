package io

import (
  "go-jvm/src/jvm/ch11/native"
  "go-jvm/src/jvm/ch11/rtda"
)

func init() {
  native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *rtda.Frame) {
  // todo
  frame.OperandStack().PushLong(0)
}
