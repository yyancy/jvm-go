package native

import "go-jvm/src/jvm/ch10/rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
  key := className + "~" + methodName + "~" + methodDescriptor
  registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
  key := className + "~" + methodName + "~" + methodDescriptor
  if method, ok := registry[key]; ok {
    return method
  }
  if methodDescriptor == "()V" && methodName == "registerNatives" {
    return emptyNativeMethod
  }
  return nil
}

// 空的nativeMethod方法
func emptyNativeMethod(frame *rtda.Frame) {
  // do nothing
}
