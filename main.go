package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jailtonjunior94/udemy-go-concorrencia/canais"
	"github.com/jailtonjunior94/udemy-go-concorrencia/waitgroup"
)

func main() {
	/* Concorrência != Paralelismo */

	go escrever("Olá mundo")
	escrever("Programando em Go!")

	/* WaitGroup */
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		waitgroup.Escrever("Aprendendo goroutines com WaitGroup")
		wg.Done()
	}()

	go func() {
		waitgroup.Escrever("Programando em Go!")
		wg.Done()
	}()

	wg.Wait()

	/* Canais */
	canal := make(chan string)
	go canais.Escrever("Olá mundo", canal)

	/* Recebendo um valor do channel */
	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	/* Recebendo um valor do channel */
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	/* Canais com Buffer */
	canalBuffer := make(chan string, 2)
	canalBuffer <- "Olá mundo!"

	mensagemBuffer := <-canalBuffer
	fmt.Println(mensagemBuffer)

}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
