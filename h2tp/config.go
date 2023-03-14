package h2tp

type MessageConfig struct {
	MaxFirstLineSize uint32
	MaxURISize       uint32
	MaxHeaderSize    uint32
	MaxHeaderCount   uint32
	MaxBodySize      uint32
}

type SocketConfig struct {
}
