package main

import "fmt"

func main() {
	arrays()
	structs()
	maps()
}

func arrays() {
	fmt.Println("----ARRAYS----")
	//los arreglos se definen con su longitud y su tipo de dato como en c
	//Si se define solamente con longitud y tipo de dato, entonces los valores se deben asignar por su index
	var array [3]string
	array[0] = "Uno"
	array[1] = "Dos"
	array[2] = "Tres"
	fmt.Println(array)

	//Para definir un array ya lleno se debe usar la sintaxis con el operador :=

	second_Array := [3]string{"Manzana", "Sandía", "Melón"}
	fmt.Println(second_Array)

	//Los slice son igual en un arreglo pero son más flexibles. Pueden tener una cantidad de objetos iniciales y un máximo de elementos, pero también puede dejarse sin limitar el máximo.

	//Se tiene referencia a un array, el segundo parámetro es con cuantos espacios arranca y el tercer parámetro es el máximo que puede contener, pero este puede estar vacío para decir que no tiene un límite fijo.
	slice := make([]string, 0)
	//Para agregar elementos hacemos un append.
	//Pude notar que podemos agregar el resultado de esta operación a un nuevo array y así tener una copia del arreglo con el elemento añadido
	slice = append(slice, "Append")
	fmt.Println(slice)
	fmt.Println(len(slice))
}

func structs() {
	fmt.Println("----STRUCTS----")
	//las structs son como las interfaces de typescript
	type Person struct {
		name string
		age  int
	}
	var new_person Person
	new_person.name = "Esteban"
	new_person.age = 21
	fmt.Println(new_person)
	//También se pueden inicializar las propiedades cuando definimos una nueva variable del tipo de la estructura

	var second_person Person = Person{name: "Luis", age: 20}
	fmt.Println(second_person)
}

func maps() {
	fmt.Println("----MAPS----")
	//Los maps son como los diccionarios de python

	a_map := make(map[string]int)
	a_map["Esteban"] = 3121032760
	a_map["Mamá"] = 3121990635
	a_map["d"] = 123
	fmt.Println(a_map)

	//Si se tratara de acceder a una llave que no esta declarada en el mapa se puede usar la asignación múltiple para saberlo

	valor, existe := a_map["d"]
	fmt.Println(valor, existe)
	//podemos elimiar elementos
	delete(a_map, "d")
	valor, existe = a_map["d"]
	fmt.Println(valor, existe)

}
