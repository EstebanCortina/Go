package main

import "fmt"

// Los tipos de datos son sencillos (por el momento)

func main() {
	// Para los strings podemos usar comillas invertidas para que imprima
	// literalmente los valores en patalla y no sean interpretados
	fmt.Print("Hola\n")
	fmt.Println(`Hola\n`)

	// Enteros y flotantess
	fmt.Println(1)
	fmt.Println(3.14)

	//booleanos
	fmt.Println(true)
	fmt.Println(false)
}
