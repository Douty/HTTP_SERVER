package pool

import (
	"bufio"
	"io"
	"sync"
)

var (
	ReadBufferPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 4096)
		}}
	ReaderPool = sync.Pool{
		New: func() interface{} {
			return bufio.NewReaderSize(nil, 4096)
		},
	}
	WriterBufferPool = sync.Pool{
		New: func() interface{} {
			return bufio.NewWriterSize(nil, 4096)
		}}
)

func GetReader(r io.Reader) *bufio.Reader {
	br := ReaderPool.Get().(*bufio.Reader)
	br.Reset(r)
	return br
}
func PutReader(br *bufio.Reader) {
	br.Reset(nil)
	ReaderPool.Put(br)
}
