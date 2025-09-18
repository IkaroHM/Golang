package main

import "fmt"

func main() {
	var Aluno string = "Ikaro"
	var nota1i float64 = 9
	var nota2i float64 = 9
	var nota3i float64 = 10
	var mediai float64 = (nota1i + nota2i + nota3i) / 3
	if mediai >= 7 {
		fmt.Printf("%s teve uma media de %.2f\n e passou", Aluno, mediai)
	} else {
		fmt.Printf("%s teve uma media de %.2f\n e nao passou", Aluno, mediai)
	}



	var Alunoa string = "Abacate"
	var nota1a float64 = 9
	var nota2a float64 = 9
	var nota3a float64 = 10
	var mediaa float64 = (nota1a + nota2a + nota3a) / 3
	if mediaa >= 7 {
		fmt.Printf("%s teve uma media de %.2f\n e passou", Alunoa, mediai)
	} else {
		fmt.Printf("%s teve uma media de %.2f\n e nao passou", Alunoa, mediai)
	}


	
	var Alunob string = "Brasil"
	var nota1b float64 = 9
	var nota2b float64 = 9
	var nota3b float64 = 10
	var mediab float64 = (nota1b + nota2b + nota3b) / 3
	if mediab >= 7 {
		fmt.Printf("%s teve uma media de %.2f\n e passou", Alunob, mediai)
	} else {
		fmt.Printf("%s teve uma media de %.2f\n e nao passou", Alunob, mediai)
	}

}
