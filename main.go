package main

import (
	"fmt"
	"github.com/Luis0714/ProyectoLL1.gitgithub.com/Luis0714/ProyectoLL1/ManejoArchivo"
)

func main(){
	ruta, err := ManejoArchivo.SeleccionarArchivo()
	archivo, err := ManejoArchivo.ObtenerArchivo(ruta)
    defer archivo.Close()
	contenido, err := ManejoArchivo.ObtenerContenido(archivo)
	gramatica, err := ManejoArchivo.DecodificarContenido(contenido)
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }

    // Decodificar el archivo JSON en una estructura de datos


    // Imprimir la estructura de datos
    fmt.Println("Terminales:", gramatica.Terminales[0])
    fmt.Println("No terminales:", gramatica.NoTerminales[0])
    fmt.Println("Expresiones:", gramatica.Expresiones["E"])
}