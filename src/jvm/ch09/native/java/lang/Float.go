package lang

import (
  "go-jvm/src/jvm/ch09/native"
  "go-jvm/src/jvm/ch09/rtda"
  "math"
)

func init() {
  native.Register("java/lang/Float",
    "floatToRawIntBits", "(F)I", floatToRawIntBits)
}

// public static native int floatToRawIntBits(float value);
func floatToRawIntBits(frame *rtda.Frame) {
  value := frame.LocalVars().GetFloat(0)
  bits := math.Float32bits(value)
  frame.OperandStack().PushInt(int32(bits))
}
