package main

import "fmt"

func soma(a int, b int) int{
	return a+b
}

func main() {
	fmt.Println("total:", soma(10, 5))
	/* arrays nao podem ter coisas a mais caso ultrapasse a capacidade dela (primeiro []).
	slices pode ter coisas adicionadas pois nao tem capacidade e nem uso predefinidos

	cap le a capacidade 
	len le o tanto de coisas
	*/
	var array [3] string = [3]string{"teste", "teste", "teste"}
	fmt.Println(array)
	fmt.Println(cap(array))
	fmt.Println(len(array))
	
	var slice []string = []string{"teste", "teste", "teste"}
	fmt.Println(slice)

	slice = append(slice, "Ikaro")

	fmt.Println(cap(slice))
	fmt.Println(len(slice))

}