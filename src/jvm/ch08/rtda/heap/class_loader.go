package heap

import (
  "fmt"
  "go-jvm/src/jvm/ch08/classfile"
  "go-jvm/src/jvm/ch08/classpath"
)

type ClassLoader struct {
  cp          *classpath.Classpath
  verboseFlag bool
  classMap    map[string]*Class // loaded classes
}

func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
  return &ClassLoader{
    cp:          cp,
    verboseFlag: verboseFlag,
    classMap:    make(map[string]*Class),
  }
}

func (self *ClassLoader) LoadClass(name string) *Class {
  if class, ok := self.classMap[name]; ok {
    return class // 类已经加载
  }
  if name[0] == '[' {
    return self.loadArrayClass(name)
  }
  return self.loadNonArrayClass(name)
}

// 加载数组类  name-> 类名   Class->数组类
func (self *ClassLoader) loadArrayClass(name string) *Class {
  class := &Class{
    accessFlags: ACC_PUBLIC, // todo
    name:        name,
    loader:      self,
    initStarted: true,
    superClass:  self.LoadClass("java/lang/Object"),
    interfaces: []*Class{
      self.LoadClass("java/lang/Cloneable"),
      self.LoadClass("java/io/Serializable"),
    },
  }
  self.classMap[name] = class
  return class
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
  data, entry := self.readClass(name)
  class := self.defineClass(data)
  link(class)
  if self.verboseFlag {
    fmt.Printf("[Loaded %s from %s]\n", name, entry)
  }
  return class
}

func link(class *Class) {
  verify(class)
  prepare(class)
}

func calcInstanceFieldSlotIds(class *Class) {
  slotId := uint(0)
  if class.superClass != nil {
    slotId = class.superClass.instanceSlotCount
  }
  for _, field := range class.fields {
    if !field.IsStatic() {
      field.slotId = slotId
      slotId++
      if field.isLongOrDouble() {
        slotId++
      }
    }
  }
  class.instanceSlotCount = slotId
}

// 给类变量分配空间和赋值
func allocAndInitStaticVars(class *Class) {
  class.staticVars = newSlots(class.staticSlotCount)
  for _, field := range class.fields {
    if field.IsStatic() && field.IsFinal() {
      initStaticFinalVar(class, field)
    }
  }
}

func initStaticFinalVar(class *Class, field *Field) {
  vars := class.staticVars
  cp := class.constantPool
  cpIndex := field.ConstValueIndex()
  slotId := field.SlotId()
  if cpIndex > 0 {
    switch field.Descriptor() {
    case "Z", "B", "C", "S", "I":
      val := cp.GetConstant(cpIndex).(int32)
      vars.SetInt(slotId, val)
    case "J":
      val := cp.GetConstant(cpIndex).(int64)
      vars.SetLong(slotId, val)
    case "F":
      val := cp.GetConstant(cpIndex).(float32)
      vars.SetFloat(slotId, val)
    case "D":
      val := cp.GetConstant(cpIndex).(float64)
      vars.SetDouble(slotId, val)
    case "Ljava/lang/String;":
      goStr := cp.GetConstant(cpIndex).(string)
      jStr := JString(class.Loader(), goStr)
      vars.SetRef(slotId, jStr)
    }
  }
}

// 计算实例字段的个数,并编号
func calcStaticFieldSlotIds(class *Class) {
  slotId := uint(0)
  for _, field := range class.fields {
    if field.IsStatic() {
      field.slotId = slotId
      slotId++
      if field.isLongOrDouble() {
        slotId++
      }
    }
  }
  class.staticSlotCount = slotId
}

// 计算静态字段的个数,并编号
func prepare(class *Class) {
  calcInstanceFieldSlotIds(class)
  calcStaticFieldSlotIds(class)
  allocAndInitStaticVars(class)
}

func verify(class *Class) {
  // todo
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
  data, entry, err := self.cp.ReadClass(name)
  if err != nil {
    panic("java.lang.ClassNotFoundException: " + name)
  }
  return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
  class := parseClass(data)
  class.loader = self
  resolveSuperClass(class)
  resolveInterfaces(class)
  self.classMap[class.name] = class
  return class
}

func parseClass(data []byte) *Class {
  cf, err := classfile.Parse(data)
  if err != nil {
    panic("java.lang.ClassFormatError")
  }
  return newClass(cf) // 见 6.1.1小节
}

func resolveSuperClass(class *Class) {
  if class.name != "java/lang/Object" {
    class.superClass = class.loader.LoadClass(class.superClassName)
  }
}

func resolveInterfaces(class *Class) {
  interfaceCount := len(class.interfaceNames)
  if interfaceCount > 0 {
    class.interfaces = make([]*Class, interfaceCount)
    for i, interfaceName := range class.interfaceNames {
      class.interfaces[i] = class.loader.LoadClass(interfaceName)
    }
  }
}
