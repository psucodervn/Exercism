package paasio

import (
	"io"
	"sync"
)

var _ ReadWriteCounter = (*rwCounter)(nil)

type rwCounter struct {
	r io.Reader
	w io.Writer

	rc, wc     int64
	rops, wops int
	sync.RWMutex
}

func (rw *rwCounter) Read(p []byte) (n int, err error) {
	if n, err = rw.r.Read(p); err != nil {
		return
	}
	rw.Lock()
	rw.rc += int64(n)
	rw.rops += 1
	rw.Unlock()
	return
}

func (rw *rwCounter) Write(p []byte) (n int, err error) {
	if n, err = rw.w.Write(p); err != nil {
		return
	}
	rw.Lock()
	rw.wc += int64(n)
	rw.wops += 1
	rw.Unlock()
	return
}

func (rw *rwCounter) ReadCount() (n int64, nops int) {
	rw.RLock()
	n, nops = rw.rc, rw.rops
	rw.RUnlock()
	return
}

func (rw *rwCounter) WriteCount() (n int64, nops int) {
	rw.RLock()
	n, nops = rw.wc, rw.wops
	rw.RUnlock()
	return
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &rwCounter{w: w}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &rwCounter{r: r}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &rwCounter{r: rw, w: rw}
}
