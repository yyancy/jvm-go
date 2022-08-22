package heap

func (self *Class) IsArray() bool {
  return self.name[0] == '['
}

func (self *Class) NewArray(count uint) *Object {
  if !self.IsArray() {
    panic("Not array class: " + self.name)
  }
  switch self.Name() {
  case "[Z":
    return &Object{class: self, data: make([]int8, count)}
  case "[B":
    return &Object{class: self, data: make([]int8, count)}
  case "[C":
    return &Object{class: self, data: make([]uint16, count)}
  case "[S":
    return &Object{class: self, data: make([]int16, count)}
  case "[I":
    return &Object{class: self, data: make([]int32, count)}
  case "[J":
    return &Object{class: self, data: make([]int64, count)}
  case "[F":
    return &Object{class: self, data: make([]float32, count)}
  case "[D":
    return &Object{class: self, data: make([]float64, count)}
  default:
    return &Object{class: self, data: make([]*Object, count)}
  }
}

// 获取组件类的对象
func (self *Class) ComponentClass() *Class {
  componentClassName := getComponentClassName(self.name)
  return self.loader.LoadClass(componentClassName)
}
