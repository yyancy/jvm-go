package heap

import "go-jvm/src/jvm/ch10/classfile"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool,
	classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
