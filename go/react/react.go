package react

const testVersion = 4

type callback *func(int)

type cell struct {
	value int
	// update func() bool
	// cpu
	callbacks []callback
	update    func() bool
	// input
	wires *reactor
}

func (c *cell) Value() int {
	return c.value
}

// input
func (c *cell) SetValue(v int) {
	if v != c.value {
		c.value = v
		c.wires.broadcast() // notify
	}
}

/*
 * Since we create input/compute in toplogical order, so is the `for` in broadcast,
 * thus make it invoked correctly.
 * But I think it's dirty to keep *reactor in input, insteadly we should only keep
 * related outputs as wires, and do toplogical traversal in broadcast.
 */
func (r *reactor) broadcast() {
	for _, out := range r.cells {
		if out.update != nil && out.update() { // input's `update` is nil
			for _, cb := range out.callbacks {
				(*cb)(out.Value())
			}
		}
	}
}

// cpu
func (c *cell) AddCallback(cb func(int)) CallbackHandle {
	c.callbacks = append(c.callbacks, &cb)
	return &cb
}

func (c *cell) RemoveCallback(h CallbackHandle) {
	newCallbacks := make([]callback, 0)
	for _, cb := range c.callbacks {
		if cb != callback(h.(*func(int))) {
			newCallbacks = append(newCallbacks, cb)
		}
	}
	c.callbacks = newCallbacks
}

type reactor struct {
	cells []*cell
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(v int) InputCell {
	i := cell{value: v, wires: r}
	r.cells = append(r.cells, &i)
	return &i
}

func (r *reactor) CreateCompute1(input Cell, cpu func(int) int) ComputeCell {
	return r.CreateComputeAux([]Cell{input}, func(cells []Cell) int {
		return cpu(cells[0].Value())
	})
}

func (r *reactor) CreateCompute2(input1, input2 Cell, cpu func(int, int) int) ComputeCell {
	return r.CreateComputeAux([]Cell{input1, input2}, func(cells []Cell) int {
		return cpu(cells[0].Value(), cells[1].Value())
	})
}

func (r *reactor) CreateComputeAux(cells []Cell, cpu func([]Cell) int) ComputeCell {
	out := cell{value: 0}
	out.update = func() bool {
		old := out.Value()
		out.value = cpu(cells)
		return out.Value() != old
	}
	r.cells = append(r.cells, &out)
	out.update() // compute right now
	return &out
}
