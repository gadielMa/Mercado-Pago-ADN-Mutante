package main

import (
	api "github.com/gadielMa/test/controller"
	"github.com/gin-gonic/gin"
)

/*

			Ejercicio realizado por Mercado Pago S.A.

		Llamaremos al json ("dna": {array de strings}) recibido por parámetro como "Dna",
	a cada string del array como "Chromosome",a cada char como "Protein" y las cadenas
	"AAAA", "TTTT", "GGGG" y "CCCC" como "Sequences".

		Realizaremos validaciones a la hora del ingreso de los datos por request.

		Transformaremos las Sequences verticales, oblicuas y oblicuas invertidas a horizontales para
		procesarlas facilmente como string con un simple strings.Contains("").

		Persistiremos en una database noSQL, debido al escalamiento horizontal y demanda de datos.

*/

// @title Mutant Go Api
// @version 1.0
// @description API Restful
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	router := gin.Default()
	api.SetupRouter(router)
	router.Run()
}
