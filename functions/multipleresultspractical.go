type Color interface {
	RGBA() (r, g, b, a uint32)
}

func Open(name string) (*File, error)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

