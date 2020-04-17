package modle

type NetInterfase interface {
	Open(host, path string) error
	Write(buf []byte) (int, error)
	Read(buf []byte) (int, error)
	Close() error
}

type LogInterfase interface {
	//printf terminal
	Printftml(v ...interface{})
	Printlntml(format string, v ...interface{})

	PrintlnErr(v ...interface{})
	PrintlnWarring(v ...interface{})
	PrintlnInfo(v ...interface{})

	PrintfErr(format string, v ...interface{})
	PrintfWarring(format string, v ...interface{})
	PrintfInfo(format string, v ...interface{})
}
