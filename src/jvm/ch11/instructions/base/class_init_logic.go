package base

import (
  "go-jvm/src/jvm/ch11/rtda"
  "go-jvm/src/jvm/ch11/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
  class.StartInit()
  scheduleClinit(thread, class)
  initSuperClass(thread, class)
}

// 准备执行cinit方法
func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
  clinit := class.GetClinitMethod()
  if clinit != nil {
    // exec <clinit>
    newFrame := thread.NewFrame(clinit)
    thread.PushFrame(newFrame)
  }
}

// 如果父类没有初始化,进行初始化
func initSuperClass(thread *rtda.Thread, class *heap.Class) {
  if !class.IsInterface() {
    superClass := class.SuperClass()
    if superClass != nil && !superClass.InitStarted() {
      InitClass(thread, superClass)
    }
  }
}
