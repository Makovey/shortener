package shortener

import (
	"fmt"

	"github.com/Makovey/shortener/internal/config"
	"github.com/Makovey/shortener/internal/logger/stdout"
	"github.com/Makovey/shortener/internal/repository/inmemory"
)

// ExampleGenerateShortURL проверка на то, что метод отдает одинаковые результаты
// для одинаковых входных параметров
func ExampleService_generateShortURL() {
	log := stdout.NewLoggerDummy()
	repo := inmemory.NewRepositoryInMemory()
	cfg := config.NewConfig(log)

	service := NewShortenerService(
		repo,
		cfg,
		log,
	)

	out1 := service.generateShortURL("https://google.com")
	fmt.Println(out1)

	out2 := service.generateShortURL("https://google.com")
	fmt.Println(out2)

	out3 := service.generateShortURL("https://facebook.com")
	fmt.Println(out3)

	// Output:
	// 99999ebcfdb78d
	// 99999ebcfdb78d
	// a023cfbf5f1c39
}
