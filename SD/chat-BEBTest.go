// Construido como parte da disciplina de Sistemas Distribuidos
// PUCRS - Escola Politecnica
// Professor: Fernando Dotti  (www.inf.pucrs.br/~fldotti)

/*
LANCAR N PROCESSOS EM SHELL's DIFERENTES, PARA CADA PROCESSO, O SEU PROPRIO ENDERECO EE O PRIMEIRO DA LISTA
go run chat.go 127.0.0.1:5001  127.0.0.1:6001    ...
go run chat.go 127.0.0.1:6001  127.0.0.1:5001    ...
go run chat.go ...  127.0.0.1:6001  127.0.0.1:5001
*/

package main

import (
	. "SD/URB"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please specify at least one address:port!")
		fmt.Println("go run chat-BEBTest.go 127.0.0.1:5001  127.0.0.1:6001 127.0.0.1:7001")
		fmt.Println("go run chat-BEBTest.go 127.0.0.1:6001  127.0.0.1:5001 127.0.0.1:7001")
		fmt.Println("go run chat-BEBTest.go 127.0.0.1:7001  127.0.0.1:6001  127.0.0.1:5001")
		return
	}

	var registro []string
	addresses := os.Args[1:]

	urb := URB_Module{
		Req: make(chan URB_Req_Message, 100),
		Ind: make(chan URB_Ind_Message, 100)}

	urb.Init(addresses[0], addresses[0:])

	// enviador de broadcasts
	go func() {

		scanner := bufio.NewScanner(os.Stdin)
		var msg string

		for {
			if scanner.Scan() {
				msg = scanner.Text()
				msg += "§" + addresses[0]
			}

			req := URB_Req_Message{
				Addresses: addresses[0:],
				Message:   msg}
			urb.Req <- req
		}
	}()

	// receptor de broadcasts
	go func() {

		for {
			in := <-urb.Ind
			message := strings.Split(in.Message, "§")
			in.From = message[1]
			registro = append(registro, strings.Split(in.Message, "§")[0])

			in.Message = message[0]

			// imprime a mensagem recebida na tela
			fmt.Printf("          Message from %v: %v\n", in.From, in.Message)

			if len(registro) == 5 {
				os.Exit(0)
			}
		}
	}()

	blq := make(chan int)
	<-blq
}
