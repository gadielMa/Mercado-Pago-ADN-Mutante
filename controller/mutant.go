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

// es la cantidad de chromosomes en el dna, como también el largo de esos chrmosomes [6x6]
const adnLength int = 6

// proteinas permitidas en Dna
var proteins = []string{"A", "T", "C", "G"}

// Mutant http GET que informa si un Dna es mutante o no
func Mutant(ctx *gin.Context) {
	var request models.Request
	validate := validator.New()

	// validaciones del json recibido
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

	// devuelve si el dna es de mutante
	ok := isMutant(request)

	// persistimos dna en mongoDb
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
func isMutant(req models.Request) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	total := 0

	// horizontal
	go numberOfSequences(req.Dna, ch1)

	// vertical
	go numberOfSequences(dnaVerticalToHorizontal(req.Dna), ch2)

	// oblicuo
	go numberOfSequences(dnaObliqueToHorizontal(req.Dna, 0), ch3)

	// oblicuo invertido
	go numberOfSequences(dnaObliqueToHorizontal(req.Dna, 1), ch4)

	total = <-ch1 + <-ch2 + <-ch3 + <-ch4

	if total < 3 {
		return false
	}

	return true
}

// cuantas sequences horizontales "XXXX" tiene un dna
func numberOfSequences(dna []string, c chan int) {
	total := 0
	for _, sequence := range dna {
		for _, protein := range proteins {
			if strings.Contains(sequence, proteinToSequence(protein)) {
				total++
			}
		}
	}
	c <- total
}

// convierte un char "A" a "AAAA"
func proteinToSequence(protein string) string {
	sequence := ""
	for i := 0; i < sequenceLength; i++ {
		sequence = sequence + protein
	}
	return sequence
}

//////////////////
// Conversiones //
//////////////////

// recibe como parametro el dna
func dnaVerticalToHorizontal(dna []string) []string {
	// creamos array de strings vacío
	var dnaHorizontal = []string{}
	for a := 0; a < adnLength; a++ {
		dnaHorizontal = append(dnaHorizontal, "")
	}

	// usamos cada string y lo transformamos en array
	for _, chromosome := range dna {
		chromosomeArray := strings.Split(chromosome, "")

		// ponemos cada elemento del nuevo array, en el array de strings vacío que creamos
		for i, protein := range chromosomeArray {
			dnaHorizontal[i] = dnaHorizontal[i] + protein
		}
	}

	return dnaHorizontal
}

// recibe como parametro el dna y una dirección oblicua
func dnaObliqueToHorizontal(dna []string, direction int) []string {
	// array de strings vacío
	var dnaHorizontal = []string{}
	var dnaHorizontalLength = adnLength*2 - 1 - ((sequenceLength - 1) * 2)
	for i := 0; i < dnaHorizontalLength; i++ {
		dnaHorizontal = append(dnaHorizontal, "")
	}

	// generamos matriz
	var matrix [adnLength][]string
	for i, chromosome := range dna {
		matrix[i] = strings.Split(chromosome, "")
	}

	// generamos un []string{} con las diagonales de la matriz, eligiendo
	// las diagonales normales o las invertidas segun el caso
	for i := 0; i < dnaHorizontalLength; i++ {
		if direction == 0 {
			dnaHorizontal[i] = diagonal(matrix, i+sequenceLength-1)
		} else {
			dnaHorizontal[i] = diagonalInverted(matrix, i+sequenceLength-1)
		}
	}

	return dnaHorizontal
}

// recime como parametro un dna en formato [][]string{} y el número de diagonal a obtener
func diagonal(matrix [adnLength][]string, diagonal int) string {
	var respuesta string
	var initial int
	row := diagonal

	// esto sucede cuando llegamos al fondo de la matriz verticalmente
	// y debemos seguir recorriendola horizontalmente para obtener todas las diagonales
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

	// cantidad de interacciones a hacer o tamaño de string a devolver
	interactions := diagonal + 1

	// esto sucede cuando llegamos al fondo de la matriz verticalmente
	// y debemos seguir recorriendola horizontalmente para obtener todas las diagonales
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
