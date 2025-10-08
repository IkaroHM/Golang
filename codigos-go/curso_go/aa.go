package main

import (
	"fmt"
)

type Cpu struct {
	 Nome string
	 chipset string
	 Marca string
}

func sla(f *Cpu) {
	f.Marca = "amtel"
}

func main() {
	ryzen := Cpu {"ryzen 5500", "am4", "amd"}
	sla(&ryzen)
	fmt.Println(ryzen.Marca)
}