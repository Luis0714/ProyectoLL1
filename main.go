package main

import (
	"fmt"
	"github.com/Luis0714/ProyectoLL1.gitgithub.com/Luis0714/ProyectoLL1/ManejoArchivo"
	"github.com/Luis0714/ProyectoLL1.gitgithub.com/Luis0714/ProyectoLL1/Logica"
	
)

func main(){
	ruta, err := ManejoArchivo.SeleccionarArchivo()
	archivo, err := ManejoArchivo.ObtenerArchivo(ruta)
    defer archivo.Close()
	contenido, err := ManejoArchivo.ObtenerContenido(archivo)
	gramatica, err := ManejoArchivo.DecodificarContenido(contenido)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    fmt.Println(Logica.Primeros(gramatica))
}