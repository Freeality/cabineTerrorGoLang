package iniciar

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

// Log representa um registro de execução do programa
type Log struct {
	Programa  string
	Intervalo time.Duration
}

// IniciarComTicker inicia o programa no intervalo de tempo da instancia
func (i *Log) IniciarComTicker() {
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

// IniciarEsperando usa CommandContext para finalizar o programa
// em um tempo determinado
func (i *Log) IniciarEsperando() {
	t0 := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), i.Intervalo)
	defer cancel()

	cmnd := exec.CommandContext(ctx, i.Programa)
	err := cmnd.Run()
	err = cmnd.Wait()

	if err != nil {
		fmt.Printf("Código de saída...%v \n", err)
		t1 := time.Now()
		fmt.Printf("Tempo de execução... %v\n", t1.Sub(t0))
	}
}
