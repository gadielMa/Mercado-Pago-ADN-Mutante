package main

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

/*

			Ejercicio de Mercado Libre S.A.

		Llamaremos al array recibido por parámetro como "Dna", a cada string del array como "Chromosome" y
	a cada char como "Protein". Y las cadenas "AAAA", "TTTT", "GGGG" y "CCCC" como "Sequence"

		Realizaremos validaciones a la hora del ingreso de los datos.

		Procesaremos en paralelo, las "Sequences" horizontales, verticales y oblicuas, posterior a una
	transformación a horizontal.

*/

const sequenceLength int = 4
const adnLength int = 6

var proteins = []string{"A", "T", "C", "G"}

func main() {
	router := gin.Default()

	router.POST("/mutant", mutant)

	router.Run()
}

type request struct {
	Dna []string `json:"dna" validate:"eq=6,dive,validate_proteins,min=6,max=6"`
}

func mutant(ctx *gin.Context) {
	var request request
	validate := validator.New()

	_ = validate.RegisterValidation("validate_proteins", validateProteins)

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(request)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ok := isMutant(request)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"mutant": "false"}) // TODO
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"mutant": "true"}) // TODO
	return
}

func validateProteins(fl validator.FieldLevel) bool {
	if regexp.MustCompile(`^[ATCG_\-.]+$`).MatchString(fl.Field().String()) {
		return true
	}
	return false
}

func isMutant(req request) bool {
	total := 0

	// Horizontal
	total = total + numberOfSequences(req.Dna)

	// Vertical
	total = total + numberOfSequences(dnaVerticalToHorizontal(req.Dna))

	// Oblicuo 1
	total = total + numberOfSequences(dnaObliqueToHorizontal(req.Dna, 0))

	// Oblicuo 2
	total = total + numberOfSequences(dnaObliqueToHorizontal(req.Dna, 1))

	if total < 3 {
		return false
	}
	return true
}

func numberOfSequences(dna []string) int {
	total := 0
	for _, sequence := range dna {
		for _, protein := range proteins {
			if strings.Contains(sequence, proteinToSequence(protein)) {
				total++
			}
		}
	}
	return total
}

func proteinToSequence(protein string) string {
	sequence := ""
	for i := 0; i < sequenceLength; i++ {
		sequence = sequence + protein
	}
	return sequence
}

func dnaVerticalToHorizontal(dna []string) []string {
	// Creamos array de strings vacío (dnaHorizontal)
	var dnaHorizontal = []string{}
	for a := 0; a < adnLength; a++ {
		dnaHorizontal = append(dnaHorizontal, "")
	}

	// Usamos cada string (chromosome) y lo transformamos en array
	for _, chromosome := range dna {
		chromosomeArray := strings.Split(chromosome, "")

		// Ponemos cada elemento del nuevo array en el array de strings vacío que creamos
		for i, protein := range chromosomeArray {
			dnaHorizontal[i] = dnaHorizontal[i] + protein
		}
	}

	return dnaHorizontal
}

func dnaObliqueToHorizontal(dna []string, dir int) []string {
	// Creamos array de strings vacío (dnaHorizontal)
	var dnaHorizontal = []string{}
	var dnaHorizontalLength = adnLength*2 - 1 - ((sequenceLength - 1) * 2)
	for i := 0; i < dnaHorizontalLength; i++ {
		dnaHorizontal = append(dnaHorizontal, "")
	}

	var matrix [adnLength][]string
	// Generamos una matrix
	for i, chromosome := range dna {
		matrix[i] = strings.Split(chromosome, "")
	}

	for i := 0; i < dnaHorizontalLength; i++ {
		if dir == 0 {
			dnaHorizontal[i] = diagonal(matrix, i+sequenceLength-1)
		} else {
			dnaHorizontal[i] = diagonalInverted(matrix, i+sequenceLength-1)
		}
	}

	return dnaHorizontal
}

func diagonal(matrix [adnLength][]string, diagonal int) string {
	var respuesta string
	var initial int
	row := diagonal

	if diagonal >= adnLength {
		difference := diagonal - adnLength + 1
		initial = difference
		diagonal = diagonal - difference
		row = adnLength + initial - 1
	}

	for column := initial; column <= diagonal; column++ {
		respuesta = respuesta + matrix[row-column][column]
	}
	return respuesta
}

func diagonalInverted(matrix [adnLength][]string, diagonal int) string {
	var response string
	var difference int

	// cantidad de chars a devolver en response
	interactions := diagonal + 1

	// difference -> esto sucede cuando llegamos al fondo de la matriz verticalmente
	// y debemos seguir recorriendola horizontalmente
	if diagonal >= adnLength {
		difference = diagonal - adnLength + 1
		interactions = adnLength*2 - 1 - diagonal
	}

	row := adnLength - difference - 1
	column := diagonal - difference

	for ; interactions > 0; interactions-- {
		response = response + matrix[row][column]
		row--
		column--
	}
	return response
}
