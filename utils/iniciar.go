package utils

import (
	"time"
	"os/exec"
	"fmt"
)

type Iniciar struct {
	Programa string
	Intervalo time.Duration
}

// inicia o programa no intervalo de tempo da instancia
func (i *Iniciar) IniciarComTicker() {
	intervaloDoTick := time.Second
	ticker := time.NewTicker(intervaloDoTick)

	cmd := exec.Command(i.Programa)
	cmd.Start()

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick em", t)
		}
	}()

	time.Sleep(i.Intervalo)
	ticker.Stop()
	cmd.Process.Kill()
}
