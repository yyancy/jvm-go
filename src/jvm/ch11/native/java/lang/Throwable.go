package lang

import (
  "fmt"
  "go-jvm/src/jvm/ch11/native"
  "go-jvm/src/jvm/ch11/rtda"
  "go-jvm/src/jvm/ch11/rtda/heap"
)

type StackTraceElement struct {
  fileName   string
  className  string
  methodName string
  lineNumber int
}

func (self *StackTraceElement) String() string {
  return fmt.Sprintf("%s.%s(%s:%d)",
    self.className, self.methodName, self.fileName, self.lineNumber)
}

func init() {
  native.Register("java/lang/Throwable", "fillInStackTrace",
    "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
func fillInStackTrace(frame *rtda.Frame) {
  this := frame.LocalVars().GetThis()
  frame.OperandStack().PushRef(this)
  stes := createStackTraceElements(this, frame.Thread())
  this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
  skip := distanceToObject(tObj.Class()) + 2
  frames := thread.GetFrames()[skip:]
  stes := make([]*StackTraceElement, len(frames))
  for i, frame := range frames {
    stes[i] = createStackTraceElement(frame)
  }
  return stes
}

// 栈顶的前几帧是异常的构造函数和方法,跳过它,计算具体跳过的个数
func distanceToObject(class *heap.Class) int {
  distance := 0
  for c := class.SuperClass(); c != nil; c = c.SuperClass() {
    distance++
  }
  return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
  method := frame.Method()
  class := method.Class()
  return &StackTraceElement{
    fileName:   class.SourceFile(),
    className:  class.JavaName(),
    methodName: method.Name(),
    lineNumber: method.GetLineNumber(frame.NextPC() - 1),
  }
}
