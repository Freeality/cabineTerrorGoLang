package utils

import (
	"time"
	"os/exec"
	"fmt"
	"log"
	"context"
)

type Iniciar struct {
	Programa string
}

func (i Iniciar) IniciarComTicker() {
	intervaloDoTick := time.Second
	intervaloTotal := time.Second * 5
	ticker := time.NewTicker(intervaloDoTick)

	cmd := exec.Command(i.Programa)
	cmd.Start()

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick em ", t)
		}
	}()

	time.Sleep(intervaloTotal)
	ticker.Stop()
	cmd.Process.Kill()
}

func (i Iniciar) IniciarSemEsperarComTimer() {
	tempoDeEspera := time.Second * 5
	t0 := time.Now()
	timer := time.NewTimer(tempoDeEspera)
	cmd := exec.Command("notepad")
	err := cmd.Start()

	if err != nil {
		fmt.Errorf("Erro tentando executar %v", err)
		return
	}

	<-timer.C
	cmd.Process.Kill()
	err = cmd.Wait()
	t1 := time.Now()
	log.Printf("Programa finalizado em %v", t1.Sub(t0))
}

func (i Iniciar) IniciarEsperando() {
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
