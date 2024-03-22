package main

import "fmt"

func main() {
	package main

import (
	"fmt"
)

func main() {
	menuprincipal()
}

func menuprincipal() {
	for {
		fmt.Println("1- Consultar tarifa")
		fmt.Println("2- Cobro de peaje")
		fmt.Println("3- Cambiar tarifa")
		fmt.Println("4- Mostrar tarifas")
		fmt.Println("5- Salir")
		fmt.Print("Ingrese una opción: ")

		var opcion int
		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			return
		}

		switch opcion {
		case 1, 2, 3, 4:
			procesarOpcion(opcion)
		case 5:
			fmt.Println("Saliendo del programa.")
			return
		default:
			fmt.Println("Opción no válida, intente nuevamente.")
		}
	}
}

func procesarOpcion(opcion int) {
	switch opcion {
	case 1:
		fmt.Println("Consultar tarifa. Selecciona el tipo de vehículo:")
	case 2:
		fmt.Println("Cobro de peaje. Selecciona el tipo de vehículo:")
	case 3:
		fmt.Println("Cambiar tarifa. Selecciona el tipo de vehículo para cambiar la tarifa:")
	case 4:
		fmt.Println("Mostrar tarifas. Aquí se mostrarían todas las tarifas disponibles.")
	default:
		fmt.Println("Opción no reconocida.")
	}
	mostrarTiposVehiculo()
}

func mostrarTiposVehiculo() {
	fmt.Println("1- Motocicleta")
	fmt.Println("2- Liviano")
	fmt.Println("3- Carga Liviana")
	fmt.Println("4- Carga Pesada")
	fmt.Println("5- Agrícola")
	fmt.Println("6- Casa Rodante")
	fmt.Println("7- Buses")
	fmt.Println("8- Regresar a menú principal")

	fmt.Print("Ingrese una opción: ")
	var seleccion int
	fmt.Scanln(&seleccion)

	if seleccion == 8 {
		return
	}
}

}
