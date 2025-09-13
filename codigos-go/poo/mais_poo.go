package main

import "fmt"

type Aluno struct{
	aluno string
	nota1 float64
	nota2 float64
	nota3 float64
	passar bool
}

func (f Aluno) passou(){
	var media float64 = (f.nota1 + f.nota2 + f.nota3) / 3
	if media >= 7 {
		fmt.Printf("%s passou pois teve uma media de %.2f\n", f.aluno, media)
		f.passar = true
	} else {
		fmt.Printf("%s nao passou pois teve uma media de %.2f\n", f.aluno, media)
		f.passar = false
	}
}

func main() {
	Alunos := []Aluno{
		{"Ikaro", 8, 9, 10, false},
		{"Jose", 7, 10, 9, false},
		{"Maria", 7, 9, 10, false},
	}
	Aluno.passou(Alunos[0])
}