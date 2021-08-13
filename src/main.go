package main

import "fmt"

func main() {
	//Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	//Suma
	x := 10
	y := 10

	result := x + y
	fmt.Println("Suma", result)

	//Resta
	result = x - y
	fmt.Println("Resta", result)

	//Multiplicación
	result = x * y
	fmt.Println("Multiplicación", result)

	//División
	result = x / y
	fmt.Println("División", result)

	//Módulo
	result = x % y
	fmt.Println("Módulo", result)

	//Incremental
	x++
	fmt.Println("Incremental", x)

	//Decremental
	x--
	fmt.Println("Decremental", x)
}
