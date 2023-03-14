package h2tp

import "io"

type Message struct {
	fl1     []byte
	fl2     []byte
	fl3     []byte
	headers Headers
	body    io.ReadWriteCloser
}

func (m *Message) XX() {
}
