package iniciar

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Log representa um registro de execução do programa
type Log struct {
	Programa    string
	Intervalo   time.Duration
	TempoMinimo time.Duration
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
		duracao := t1.Sub(t0)
		fmt.Printf("Tempo de execução... %v\n", duracao)
		fmt.Printf("data: %s, hora: %s\n", tempoParaDataString(t1), tempoParaHora(t1))

		home := os.ExpandEnv("${HOMEPATH}")

		if len(home) > 0 && duracao > i.TempoMinimo {
			criaArquivo(duracao, t1, fmt.Sprintf("%s\\logs", home))
		}
	}
}

// listaEnv lista todas as variáveis do sistema
func listaEnv() {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

// check emite panic se houver erro
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// criarArquivo cria um arquivo com nome igual a hora
// no formato hh-mm-ss na pasta informada, $HOME\logs,
// com o time.Time informado
func criaArquivo(duracao time.Duration, tempo time.Time, path string) {
	pathCompleto := fmt.Sprintf("%s\\%s\\%s.txt", path, tempoParaDataString(tempo), tempoParaHora(tempo))
	texto := []byte(fmt.Sprintf("%v", duracao))

	fmt.Printf("crindo arquivos em %s", path)
	errPath := os.MkdirAll(fmt.Sprintf("%s\\%s", path, tempoParaDataString(tempo)), 0777)
	check(errPath)

	f, err := os.Create(pathCompleto)
	check(err)

	defer f.Close()
	_, err = f.Write(texto)
	f.Sync()
}

// tempoParaDataString solicita um time.Time e retorna uma string
// com a data no formato dd-mm-aaaa
func tempoParaDataString(tempo time.Time) string {
	dia := tempo.Day()
	mes := tempo.Month()
	ano := tempo.Year()

	return fmt.Sprintf("%02d-%02d-%d", dia, mes, ano)
}

// tempoParaHora solicita um time.Time e retorna uma string
// com a hora no formato hh-mm
func tempoParaHora(tempo time.Time) string {
	hora := tempo.Hour()
	minuto := tempo.Minute()
	segundo := tempo.Second()

	return fmt.Sprintf("%d-%d-%d", hora, minuto, segundo)
}
