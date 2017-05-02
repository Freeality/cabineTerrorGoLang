package main

import (
	"os"
	"time"

	"fmt"

	u "github.com/freeality/cabineTerrorGoLang/log"
)

// O usuário deve poder utilizar o jogo por 5 minutos.
// após esse tempo o jogo deve ser encerrado
// para iniciar o jogo demora-se aproximadamente 3 minutos.
// considera-se que o usuário utilizou o jogo a partir de 2 minutos.
// cada vez que o usuário utiliza o jogo um registro deve ser armazenado
// em um arquivo para consulta posterior.
func main() {
	nomeDoPrograma := os.Args[1]

	if len(nomeDoPrograma) == 0 {
		fmt.Printf("O nome do programa deve ser especificado")
		return
	}

	// temporariamente está programado para rodar por 5 segundos
	i := u.Log{Programa: nomeDoPrograma, Intervalo: time.Second * 5, TempoMinimo: time.Second * 2}
	i.IniciarEsperando()
}
