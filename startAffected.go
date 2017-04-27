package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	iniciarAffected()
}

func iniciarAffected() {

	tempoEspera := 5 * time.Second
	timer := time.NewTimer(tempoEspera)

	go func() {
		for t := range timer.C {
			intervalo := time.Since(t)
			//if intervalo > 1 {
			fmt.Printf("Contando... %s\n", intervalo)
			//}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), tempoEspera)
	defer cancel()

	cmnd := exec.CommandContext(ctx, "Affected v1.62.exe", "arg")
	err := cmnd.Run()

	if err != nil {
		fmt.Printf("Código de saída...%v \n", err)
	}
}
