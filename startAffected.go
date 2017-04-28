package main

import (
	u "github.com/freeality/cabineTerrorGoLang/utils"
	"time"
)

// O usuário deve poder utilizar o jogo por 5 minutos.
// após esse tempo o jogo deve ser encerrado
// para iniciar o jogo demora-se aproximadamente 3 minutos.
// considera-se que o usuário utilizou o jogo a partir de 2 minutos.
// cada vez que o usuário utiliza o jogo um registro deve ser armazenado
// em um arquivo para consulta posterior.
func main() {
	i := u.Iniciar{Programa: "notepad", Intervalo:time.Second * 5}
	i.IniciarComTicker()
}
