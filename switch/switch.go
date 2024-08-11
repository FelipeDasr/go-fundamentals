package main

func getDayNameInPTBr(day int) string {
	switch day {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func getMonthNameInPTBr(month int) string {
	switch {
	case month == 1:
		return "Janeiro"
	case month == 2:
		return "Fevereiro"
	case month == 3:
		return "Março"
	case month == 4:
		return "Abril"
	case month == 5:
		return "Maio"
	case month == 6:
		return "Junho"
	case month == 7:
		return "Julho"
	case month == 8:
		return "Agosto"
	case month == 9:
		return "Setembro"
	case month == 10:
		return "Outubro"
	case month == 11:
		return "Novembro"
	case month == 12:
		return "Dezembro"
	default:
		return "Mês inválido"
	}
}

func main() {
	day := 3
	println(getDayNameInPTBr(day))
}