package controller

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gadielMa/test/db"
	"github.com/gadielMa/test/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// largo de las cadenas "AAAA", ...
const sequenceLength int = 4

// cantidad de chromosomes en el dna, como también el largo de esos chrmosomes [6x6]
const adnLength int = 6

// proteinas permitidas en Dna
var proteins = []string{"A", "T", "C", "G"}

// Mutant godoc
// @Description recibe un json y devuelve si es un mutante o no
// @Accept json
// @Produce json
// @Param Request body models.Mutant true "Datos necesarios para dar de alta un humano"
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /mutant [post]
func Mutant(ctx *gin.Context) {
	var request models.Mutant

	validate := validator.New()
	_ = validate.RegisterValidation("validate_proteins", validateProteins)

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ok := isMutant(request)

	err = db.InsertDna(request, ok)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, gin.H{"error": "No se pudo insertar en base de datos"})
		return
	}

	if !ok {
		ctx.JSON(http.StatusForbidden, nil)
		return
	}

	ctx.JSON(http.StatusOK, nil)
	return
}

// validamos los caracteres A, T, C y G
func validateProteins(fl validator.FieldLevel) bool {
	if regexp.MustCompile(`^[ATCG_\-.]+$`).MatchString(fl.Field().String()) {
		return true
	}
	return false
}

// procesamos las 4 Sequences en paralelo
func isMutant(mutant models.Mutant) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	var sum int

	go numberOfSequences(mutant.Dna, ch1)
	go numberOfSequences(dnaVerticalToHorizontal(mutant.Dna), ch2)
	go numberOfSequences(dnaObliqueToHorizontal(mutant.Dna, 0), ch3)
	go numberOfSequences(dnaObliqueToHorizontal(mutant.Dna, 1), ch4)

	sum = <-ch1 + <-ch2 + <-ch3 + <-ch4

	if sum < 2 {
		return false
	}

	return true
}

// cuantas sequences horizontales "XXXX" tiene un dna
func numberOfSequences(dna []string, c chan int) int {
	var total int
	for _, sequence := range dna {
		for _, protein := range proteins {
			if strings.Contains(sequence, proteinToSequence(protein)) {
				total++
			}
		}
	}
	c <- total
	return total
}

// convierte un char "A" a "AAAA"
func proteinToSequence(protein string) string {
	var sequence string
	for i := 0; i < sequenceLength; i++ {
		sequence = sequence + protein
	}
	return sequence
}

// recibe como parametro el dna vertial y lo devuelve horizontal
func dnaVerticalToHorizontal(dna []string) []string {
	dnaHorizontal := generateArrayOfStrings(adnLength)

	for _, chromosome := range dna {
		chromosomeArray := strings.Split(chromosome, "")

		for i, protein := range chromosomeArray {
			dnaHorizontal[i] = dnaHorizontal[i] + protein
		}
	}

	return dnaHorizontal
}

// recibe como parametro el dna y una dirección oblicua normal o invertida
func dnaObliqueToHorizontal(dna []string, direction int) []string {
	dnaHorizontalLength := adnLength*2 - 1 - ((sequenceLength - 1) * 2)
	dnaHorizontal := generateArrayOfStrings(dnaHorizontalLength)

	var matrix [adnLength][]string
	for i, chromosome := range dna {
		matrix[i] = strings.Split(chromosome, "")
	}

	for i := 0; i < dnaHorizontalLength; i++ {
		if direction == 0 {
			dnaHorizontal[i] = diagonal(matrix, i+sequenceLength-1)
		} else {
			dnaHorizontal[i] = diagonalInverted(matrix, i+sequenceLength-1)
		}
	}

	return dnaHorizontal
}

func generateArrayOfStrings(size int) []string {
	var dnaHorizontal = []string{}
	for a := 0; a < size; a++ {
		dnaHorizontal = append(dnaHorizontal, "")
	}
	return dnaHorizontal
}

// recime como parametro un dna en formato [][]string{} y el número de diagonal a obtener
func diagonal(matrix [adnLength][]string, diagonal int) string {
	var respuesta string
	var initial int
	row := diagonal

	// llegamos al fondo de la matriz verticalmente
	// y debemos recorrerla horizontalmente
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

// recime como parametro un dna en formato [][]string{} y el número de diagonal a obtener
func diagonalInverted(matrix [adnLength][]string, diagonal int) string {
	var response string
	var difference int

	interactions := diagonal + 1

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
