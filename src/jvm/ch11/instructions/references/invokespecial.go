package references

import (
  "go-jvm/src/jvm/ch11/instructions/base"
  "go-jvm/src/jvm/ch11/rtda"
  "go-jvm/src/jvm/ch11/rtda/heap"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
  currentClass := frame.Method().Class()
  cp := currentClass.ConstantPool()
  methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
  resolvedClass := methodRef.ResolvedClass()
  resolvedMethod := methodRef.ResolvedMethod()

  if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
    panic("java.lang.NoSuchMethodError")
  }
  if resolvedMethod.IsStatic() {
    panic("java.lang.IncompatibleClassChangeError")
  }
  // 得到当前引用
  ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
  if ref == nil {
    panic("java.lang.NullPointerException")
  }
  // 上面的判断确保protected方法只能被声明该方法的类或子类调用。
  if resolvedMethod.IsProtected() &&
      resolvedMethod.Class().IsSuperClassOf(currentClass) &&
      resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
      ref.Class() != currentClass &&
      !ref.Class().IsSubClassOf(currentClass) {
    panic("java.lang.IllegalAccessError")
  }

  methodToBeInvoked := resolvedMethod
  if currentClass.IsSuper() &&
      resolvedClass.IsSuperClassOf(currentClass) &&
      resolvedMethod.Name() != "<init>" {
    methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
      methodRef.Name(), methodRef.Descriptor())
  }

  if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
    panic("java.lang.AbstractMethodError")
  }
  base.InvokeMethod(frame, methodToBeInvoked)
}
