package operadora

import (
	"cursoDeGo/variaveis/pacotes/prefixo"
	"strconv"
)

//NomeOperadora nome da operadora
var NomeOperadora = "TIM"

//PrefixoDaCapitalOperadora campos juntos
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
