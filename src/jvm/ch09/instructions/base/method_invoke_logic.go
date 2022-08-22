package base

import (
  "go-jvm/src/jvm/ch09/rtda"
  "go-jvm/src/jvm/ch09/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
  thread := invokerFrame.Thread()
  newFrame := thread.NewFrame(method)
  thread.PushFrame(newFrame)

  argSlotCount := int(method.ArgSlotCount())
  if argSlotCount > 0 {
    for i := argSlotCount - 1; i >= 0; i-- {
      slot := invokerFrame.OperandStack().PopSlot()
      newFrame.LocalVars().SetSlot(uint(i), slot)
    }
  }
}
