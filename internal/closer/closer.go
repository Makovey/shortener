package closer

import (
	"io"
)

type Closer struct {
	closers []io.Closer
}

func NewCloser() *Closer {
	return &Closer{make([]io.Closer, 0)}
}

func (c *Closer) Add(closer io.Closer) {
	c.closers = append(c.closers, closer)
}

func (c *Closer) CloseAll() error {
	for _, closer := range c.closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}

	return nil
}
