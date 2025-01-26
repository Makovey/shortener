package utils

import (
	"context"
	"sync"
)

// Generator принимает массив данных и записывает их в канал
func Generator[T any](ctx context.Context, input []T) chan T {
	inputCh := make(chan T)
	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			case inputCh <- data:
			case <-ctx.Done():
				return
			}
		}
	}()

	return inputCh
}

// FanOut распараллеливает работу
// Принимает количество воркеров и замыкание/работу, которую необходимо выполнить
// Возвращает слайс каналов с результатами работ
func FanOut[T any](ctx context.Context, numWorkers int, work func(ctx context.Context) chan T) []chan T {
	channels := make([]chan T, numWorkers)

	for i := 0; i < numWorkers; i++ {
		channels[i] = work(ctx)
	}

	return channels
}

// FanIn собирает резульаты работы в один канал из списка каналов, работает в паре FanOut
// Принимает на вход список каналов, в котором результаты работы
// Отдает один канал с результатом
func FanIn[T any](ctx context.Context, bufSize int, resultChs ...chan T) chan T {
	finalCh := make(chan T, bufSize)

	var wg sync.WaitGroup

	for _, ch := range resultChs {
		chClosure := ch
		wg.Add(1)

		go func() {
			defer wg.Done()

			for data := range chClosure {
				select {
				case <-ctx.Done():
					return
				case finalCh <- data:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(finalCh)
	}()

	return finalCh
}
