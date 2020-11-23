package canais

import (
	"time"
)

// Escrever is a function
func Escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		/* Enviando um valor para o channel */
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
