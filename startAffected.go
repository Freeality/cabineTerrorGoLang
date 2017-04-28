package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// iniciarAffected()
	iniciarAffected()
}

func iniciarAffectedSemEsperar() {

}

func iniciarAffected() {
	t0 := time.Now()
	tempoEspera := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), tempoEspera)
	defer cancel()

	cmnd := exec.CommandContext(ctx, "notepad")
	err := cmnd.Run()

	if err != nil {
		fmt.Printf("Código de saída...%v \n", err)
		t1 := time.Now()
		fmt.Printf("Tempo de execução... %v\n", t1.Sub(t0))
	}
}
