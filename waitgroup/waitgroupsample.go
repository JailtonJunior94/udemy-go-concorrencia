package waitgroup

import (
	"fmt"
	"time"
)

// Escrever is a function
func Escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
