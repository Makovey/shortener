package shortener

import (
	"fmt"
)

// ExampleGenerateShortURL проверка на то, что метод отдает одинаковые результаты
// для одинаковых входных параметров
func ExampleService_generateShortURL() {
	out1 := generateShortURL("https://google.com")
	fmt.Println(out1)

	out2 := generateShortURL("https://google.com")
	fmt.Println(out2)

	out3 := generateShortURL("https://facebook.com")
	fmt.Println(out3)

	// Output:
	// 99999ebcfdb78d
	// 99999ebcfdb78d
	// a023cfbf5f1c39
}
