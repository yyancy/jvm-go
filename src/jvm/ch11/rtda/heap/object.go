package heap

type Object struct {
  class *Class
  //fields Slots
  // 普通对象的类型还是Slots
  data  interface{}
  extra interface{} // 用来记录Object实例额外的信息,可以是抛出异常时的堆栈信息
}

func (self *Object) Data() interface{} {
  return self.data
}

func (self *Object) SetData(data interface{}) {
  self.data = data
}

func newObject(class *Class) *Object {
  return &Object{
    class: class,
    data:  newSlots(class.instanceSlotCount),
  }
}

func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
  field := self.class.getField(name, descriptor, false)
  slots := self.data.(Slots)
  slots.SetRef(field.slotId, ref)
}
func (self *Object) GetRefVar(name, descriptor string) *Object {
  field := self.class.getField(name, descriptor, false)
  slots := self.data.(Slots)
  return slots.GetRef(field.slotId)
}

func (self *Object) GetIntVar(name, descriptor string) int32 {
  field := self.class.getField(name, descriptor, false)
  slots := self.data.(Slots)
  return slots.GetInt(field.slotId)
}

func (self *Object) SetIntVar(name, descriptor string, val int32) {
  field := self.class.getField(name, descriptor, false)
  slots := self.data.(Slots)
  slots.SetInt(field.slotId, val)
}

// getters
func (self *Object) Class() *Class {
  return self.class
}
func (self *Object) Fields() Slots {
  return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
  return class.IsAssignableFrom(self.class)
}

func (self *Object) Extra() interface{} {
  return self.extra
}

func (self *Object) SetExtra(v interface{}) {
  self.extra = v
}
