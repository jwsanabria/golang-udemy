package main

import "fmt"

func main() {
	var numero1, numero2 int
	fmt.Println("Por favor ingrese el primer numero: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese el segundo numero:")
	fmt.Scanln(&numero2)
	fmt.Printf("El cociente de los numeros es %d / %d es: %d \n", numero1, numero2, numero1/numero2)
	fmt.Printf("El residuo de los numeros es %d mod %d es: %d \n", numero1, numero2, numero1%numero2)
}
