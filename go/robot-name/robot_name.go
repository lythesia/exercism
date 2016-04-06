package robotname

import (
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

func gen() string {
	arr := make([]byte, 5)
	c := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {
		arr[c] = byte(65 + rand.Intn(26))
		c++
	}
	for i := 0; i < 3; i++ {
		arr[c] = byte(48 + rand.Intn(10))
		c++
	}
	return string(arr)
}

func (r *Robot) Name() string {
	if r.name == "" {
		r.name = gen()
	}
	return r.name
}

func (r *Robot) Reset() {
	r.name = gen()
}
