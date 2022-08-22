package heap

// 查找指定类中是否存在指定描述的方法
func LookupMethodInClass(class *Class, name, descriptor string) *Method {
  for c := class; c != nil; c = c.superClass {
    for _, method := range c.methods {
      if method.name == name && method.descriptor == descriptor {
        return method
      }
    }
  }
  return nil
}
func lookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
  for _, iface := range ifaces {
    for _, method := range iface.methods {
      if method.name == name && method.descriptor == descriptor {
        return method
      }
    }
    method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
    if method != nil {
      return method
    }
  }
  return nil
}
