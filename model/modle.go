package modle

type NetInterfase interface {
	Open(host, path string) error
	Write(buf []byte) (int, error)
	Read(buf []byte) (int, error)
	Close() error
}
