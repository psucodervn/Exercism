package react

import (
	"sync"
)

type Callback func(int)

type cell struct {
	cbIdx     int
	val       int
	callbacks map[int]Callback
	sync.RWMutex
}

func (c *cell) Value() int {
	return c.val
}

func (c *cell) SetValue(val int) {
	if c.val != val {
		c.val = val
		for _, cb := range c.callbacks {
			cb(val)
		}
	}
}

type canceler struct {
	Callback func()
}

func (cc *canceler) Cancel() {
	cc.Callback()
}

func (c *cell) AddCallback(cb func(val int)) Canceler {
	c.cbIdx++
	c.callbacks[c.cbIdx] = cb
	idx := c.cbIdx
	return &canceler{Callback: func() { delete(c.callbacks, idx) }}
}

func makeCell(val int) *cell {
	return &cell{val: val, callbacks: make(map[int]Callback)}
}

type reactor struct {
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(val int) InputCell {
	return makeCell(val)
}

func (r *reactor) CreateCompute1(inputCell Cell, fn func(int) int) ComputeCell {
	cc := inputCell.(*cell)
	c := makeCell(fn(cc.Value()))
	cc.AddCallback(func(v int) { c.SetValue(fn(v)) })
	return c
}

func (r *reactor) CreateCompute2(c1 Cell, c2 Cell, fn func(int, int) int) ComputeCell {
	cc1 := c1.(*cell)
	cc2 := c2.(*cell)
	c := makeCell(fn(cc1.Value(), cc2.Value()))
	cb := func(v int) {
		newVal := fn(cc1.Value(), cc2.Value())
		if c.Value() != newVal {
			c.SetValue(newVal)
		}
	}
	cc1.AddCallback(cb)
	cc2.AddCallback(cb)
	return c
}
