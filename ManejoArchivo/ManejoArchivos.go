package ManejoArchivo

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"github.com/Luis0714/ProyectoLL1.gitgithub.com/Luis0714/ProyectoLL1/Entidades"
	"github.com/sqweek/dialog"
)

func SeleccionarArchivo() (string, error) {
	file, err := dialog.File().Title("Seleccionar archivo").Filter("Todos los archivos", "*.*").Load()
	if err != nil {
		return "", err
	}
	return filepath.ToSlash(file), nil
}

func ObtenerArchivo(ruta string) (*os.File, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	return archivo, nil
}

func ObtenerContenido(archivo *os.File) ([]byte, error) {
	contenido, err := io.ReadAll(archivo)
	if err != nil {
		return nil, err
	}
	return contenido, nil
}

func DecodificarContenido(contenido []byte) (Entidades.Gramatica, error) {
	var gramatica Entidades.Gramatica
	err := json.Unmarshal(contenido, &gramatica)
	if err != nil {
		return Entidades.Gramatica{}, err
	}
	return gramatica, nil
}
