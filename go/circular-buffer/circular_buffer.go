package circular

import (
	"fmt"
)

const testVersion = 3

type Buffer struct {
	buf   []byte
	cap   int
	front int
	rear  int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{make([]byte, size+1), size + 1, 0, 0}
}

func (b *Buffer) empty() bool {
	return b.front == b.rear
}

func (b *Buffer) full() bool {
	return b.front == (b.rear+1)%b.cap
}

func (b *Buffer) ReadByte() (byte, error) {
	var c byte
	if b.empty() {
		return c, fmt.Errorf("buffer empty")
	}
	c = b.buf[b.front]
	b.front = (b.front + 1) % b.cap
	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.full() {
		return fmt.Errorf("buffer full")
	}
	b.buf[b.rear] = c
	b.rear = (b.rear + 1) % b.cap
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.full() {
		b.ReadByte()
	}
	b.WriteByte(c)
}

func (b *Buffer) Reset() {
	b.front, b.rear = 0, 0
}
