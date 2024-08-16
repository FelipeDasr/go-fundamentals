package main

func main() {
	week := map[string]string{}

	week["Sunday"] = "Domingo"
	week["Monday"] = "Segunda-feira"
	week["Tuesday"] = "Terça-feira"
	week["Wednesday"] = "Quarta-feira"
	week["Thursday"] = "Quinta-feira"
	week["Friday"] = "Sexta-feira"
	week["Saturday"] = "Sábado"

	for key, value := range week {
		println(key, ":", value)
	}
}