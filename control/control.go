package main

import "fmt"

func main() {
	//Las estructuras de control como if o for son muy similares a python y c/js combinadas todas

	for i := 1; i <= 10; i++ {
		fmt.Println(i, ".-")
	}

	j := 1
	for j < 3 {
		fmt.Println(j)
		j++
	}

	arr := [2]string{"Elemento 1", "Elemento 2"}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	//condicionales son lo mismo de siempre
	control := 21
	if control < 10 {
		fmt.Println("Menor que 10")
	} else if control <= 20 {
		fmt.Println("Valor entre 10 y 20")
	} else {
		fmt.Println("Mayor a 20")
	}

	//Algo interesante es que podemos declarar e inicializar una variable en un if

	if x := 2; x > 1 {
		fmt.Println("X mayor que 1")
	}

}
