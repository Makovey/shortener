// Package closer, содержит механизмы закрытия ресурсов.
// Такие как БД, файлы и тд.
package closer

import (
	"io"
)

// Closer содержит массив closer'ы и методы работы с ним
type Closer struct {
	closers []io.Closer // сущности, которые реализуют интрефейс io.Closer
}

// NewCloser конструктор Closer и создает пустой массив closer'ов
func NewCloser() *Closer {
	return &Closer{make([]io.Closer, 0)}
}

// Add добавляет сущность реализующую интерфейс io.Closer к общему списку
func (c *Closer) Add(closer io.Closer) {
	c.closers = append(c.closers, closer)
}

// CloseAll закрывает все добавленные closer'ы
func (c *Closer) CloseAll() error {
	for _, closer := range c.closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}

	return nil
}
