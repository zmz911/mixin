package network

const (
	TransportMessageVersion    = 2
	TransportMessageMaxSize    = 32 * 1024 * 1024
	TransportMessageHeaderSize = 6

	TransportCompressionGzip   = 1
	TransportCompressionZstd   = 2
	TransportCompressionMethod = TransportCompressionZstd
)

type TransportMessage struct {
	Version     uint8
	Compression uint8
	Size        uint32
	Data        []byte
}

type Client interface {
	Receive() ([]byte, error)
	Send([]byte) error
	Close() error
}

type Transport interface {
	Listen() error
	Dial() (Client, error)
	Accept() (Client, error)
}
