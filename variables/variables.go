package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Declaramos vaiables con var
	//Go es un lenguaje tipado pero también puede ser inferido

	var a int = 10
	fmt.Println(a)

	var b = "String"
	fmt.Println(b)

	//Las variables deben definirse y después asignarse como en C
	//Pero se puede hacer ambas con el operador := SOLO adentro de funciones

	c := 12.5
	fmt.Println(c)

	//También se pueden definir en una misma línea

	var e, f, g = true, 10, "cadena"

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	//Constantes

	const name string = "Esteban"
	fmt.Println(name)

	//Podemos declarar varias constantes

	const (
		c1 = 10
		c2 = 12
		c3 //en este caso Go toma como valor la declaración anterior
	)
	//Y también podemos imprimir varias cosas a la vez
	fmt.Println(c1, c2, c3)

	//Estas constantes se llaman de expresión y significa que su valor
	//será definido en tiempo de compilación.
	//Solo pueden usarse con valores y caracteres constantes, es decir
	//se puede definir su valor con otros valores ya definidos
	const result int = 5 * 2 //El número 5 nunca va a cambiar de valor, ni el dos, ni el signo de multiplicaciòn
	fmt.Println(result)
	//Ejemplo:
	const pi float64 = 3.14

	const new_result = pi * 10 * 10
	fmt.Println(new_result)

	//Convertir variables, castear
	var a_int int = 10
	var a_float float64 = float64(a_int)
	fmt.Println(a_float)

	//Parsear
	var a_string = "123"
	var new_int, _ = strconv.ParseFloat(a_string, 64) //el guión bajo es para indicar que no queremos manejar el error, ya que este métido nos regresa un float o un error entonces el error estaría almacenado en la segunda variable pero no queremos eso por el momento
	fmt.Println(new_int)

	//Los operadores lógimos son los mismos que en los otros lenguajes

}
