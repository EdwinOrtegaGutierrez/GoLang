package main

import "fmt"

func main() {

	var a, b int = 0, 0
	var operacion string = ""

	fmt.Println("Calculadora: Suma y resta")

	fmt.Println("\nEscoge una opcion")
	fmt.Scanln(&operacion)

	switch operacion {
	case "resta":
		fmt.Println("\nIntroduce los numeros: ")
		fmt.Scanln(&a, &b)
		fmt.Println("El resultado de la suma es: ", a-b)
	case "suma":
		fmt.Println("Introduce los numeros: ")
		fmt.Scanln(&a, &b)
		fmt.Println("El resultado de la suma es: ", a+b)
	}
}
