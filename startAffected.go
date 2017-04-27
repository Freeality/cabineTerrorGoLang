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
	t0 := time.Now()
	tempoEspera := 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), tempoEspera)
	defer cancel()

	cmnd := exec.CommandContext(ctx, "Affected v1.62.exe", "arg")
	err := cmnd.Run()

	if err != nil {
		fmt.Printf("Código de saída...%v \n", err)
		t1 := time.Now()
		fmt.Printf("Tempo de execução... %v\n", t1.Sub(t0))
	}
}
