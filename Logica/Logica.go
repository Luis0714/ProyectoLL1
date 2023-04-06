package Logica

import(
	"fmt"
	"github.com/Luis0714/ProyectoLL1.gitgithub.com/Luis0714/ProyectoLL1/Entidades"
)

func Primeros(gramatica Entidades.Gramatica) (string, error) {
	
	var primeros string
	for i := 0; i < len(gramatica.NoTerminales); i++ {
		NoTerminal := gramatica.NoTerminales[i]
		
		primer, err := ObtenerPrimerosNoTer(NoTerminal,gramatica)
		if err != nil{
			fmt.Println("Error")
		}
		primeros += primer
	}
	return primeros,nil
}

func ObtenerPrimerosNoTer(NoTerminal string, gramatica Entidades.Gramatica) (string,error){
	Prodiccion := gramatica.Expresiones[NoTerminal]
	var primeros string
	var uno string
	var dos string
	
		if len(Prodiccion[1]) > 0 {
			uno = Prodiccion[0][0]
			dos = Prodiccion[1][0]

			if Contiene(uno, gramatica.NoTerminales) {
				uno , err:= ObtenerPrimerosNoTer(uno, gramatica)
				if err != nil{
					fmt.Println("Error")
				}
				uno = uno
			}
			
			if Contiene(dos, gramatica.NoTerminales) {
				dos , err:= ObtenerPrimerosNoTer(dos, gramatica)
				if err != nil{
					fmt.Println("Error")
				}
				dos = dos
			}
			primeros +=  "Prim("+ NoTerminal +") = "+ "  { " +uno+" , "+dos+" }"+"\n"
		
		}else {
			if Contiene(uno, gramatica.NoTerminales) {
				uno , err:= ObtenerPrimerosNoTer(uno, gramatica)
				if err != nil{
					fmt.Println("Error")
				}
				uno = uno
			}
			primeros +=  "Prim("+ NoTerminal +") = "+ "  { " +uno+" }"+"\n"
		}
		return primeros,nil
}

func Contiene(s string, lista []string) bool {
    for _, v := range lista {
        if v == s {
            return true
        }
    }
    return false
}