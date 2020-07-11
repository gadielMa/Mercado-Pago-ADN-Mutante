package main

import (
	"github.com/gadielMa/test/controller"
	"github.com/gin-gonic/gin"
)

/*

			Ejercicio de Mercado Libre S.A.

		Llamaremos al json ("dna": {array de strings}) recibido por parámetro como "Dna",
	a cada string del array como "Chromosome",a cada char como "Protein" y las cadenas
	"AAAA", "TTTT", "GGGG" y "CCCC" como "Sequences".

		Realizaremos validaciones a la hora del ingreso de los datos por request.

		Transformaremos las Sequences verticales, oblicuas y oblicuas invertidas a horizontales para
		procesarlas facilmente como string con un simple strings.Contains("").

		Persistiremos en una database noSQL, debido al escalamiento horizontal y demanda de datos.

*/

func main() {
	router := gin.Default()

	router.POST("/mutant", controller.Mutant)

	router.Run()
}
