package rtda

import "go-jvm/src/jvm/ch11/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
