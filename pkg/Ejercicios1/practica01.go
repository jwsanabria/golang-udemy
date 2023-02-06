package main

import "fmt"

func main() {
	var numero1, numero2 int
	fmt.Println("Por favor ingrese el primer numero: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese el segundo numero:")
	fmt.Scanln(&numero2)
	fmt.Printf("La suma de los numeros %d + %d es: %d \n", numero1, numero2, numero1+numero2)
}
