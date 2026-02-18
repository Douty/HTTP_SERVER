package util

import "sync"

var (
	ReadBufferPool  = sync.Pool{New: func() interface{} { return make([]byte, 4096) }}
	WriteBufferPool = sync.Pool{New: func() interface{} { return make([]byte, 4096) }}
)
