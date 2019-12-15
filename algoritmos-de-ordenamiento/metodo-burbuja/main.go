package main

import "fmt"

func tomarDatosDelUsuario() (numeros []int) {

	var numero int
	var opcion string
	var ingresarNumero bool = true

	for ingresarNumero {

		fmt.Println("**********************************************")
		fmt.Print(">>             Ingrese Un Número                : ")
		fmt.Scan(&numero)
		fmt.Println("**********************************************\n")

		numeros = append(numeros, int(numero))

		fmt.Println("Desea Ingresar otro número ? (s/n) :")
		fmt.Scan(&opcion)

		if opcion != "s" {
			break
		}

	}

	return

}

func burbujaAscendente(numeros *[]int) []int {

	var (
		guardarPosActual int
		datos            []int = *numeros
		count            int   = (len(datos) - 1)
	)

	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if datos[j] > datos[j+1] {
				guardarPosActual = datos[j]
				datos[j] = datos[j+1]
				datos[j+1] = guardarPosActual
			}
		}
	}

	return datos
}

func burbujaDescendente(numeros []int) (nuevoArreglo []int) {

	var (
		index int
		count int = (len(numeros) - 1)
	)

	for count >= index {
		nuevoArreglo = append(nuevoArreglo, numeros[count])
		count--
	}

	return
}

func main() {
	numeros := tomarDatosDelUsuario()
	ordenados := burbujaAscendente(&numeros)
	fmt.Println("Datos ordenados de forma ascendente : ", ordenados)
	descendente := burbujaDescendente(ordenados)
	fmt.Println("Datos ordenados de forma descendente : ", descendente)
}
