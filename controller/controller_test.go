package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gadielMa/test/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestMutant(t *testing.T) {
	router := gin.Default()
	SetupRouter(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/mutant", strings.NewReader(`{
		"dna": [
				"TTGCGA",
				"CAGTGC",
				"TTATGT",
				"AAAAGG",
				"CTCCTA",
				"TCACTG"
			]
		}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestMutant2(t *testing.T) {
	router := gin.Default()
	SetupRouter(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/mutant", strings.NewReader(`{
		"dna": [
				"XXXXXX",
				"CAGTGC",
				"TTATGT",
				"AATAGG",
				"CTCCTA",
				"TCACTG"
			]
		}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestMutant3(t *testing.T) {
	router := gin.Default()
	SetupRouter(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/mutant", strings.NewReader(`{
		"dna": [
				"TTGCGA",
				"CAGTGC",
				"TTATGT",
				"ATAAGG",
				"CTCCTA",
				"TCACTG"
			]
		}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 403, w.Code)
}

func TestStats(t *testing.T) {
	router := gin.Default()
	SetupRouter(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/stats", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
func TestIsMutant(t *testing.T) {
	mutant := models.Mutant{Dna: []string{"ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF"}}
	assert.Equal(t, true, isMutant(mutant))

}

func TestIsMutant2(t *testing.T) {
	mutant := models.Mutant{Dna: []string{"QWERTY", "ASDFGH", "ZXCVBN", "ABCDEF", "ABCDEF", "ABCDEF"}}
	assert.Equal(t, false, isMutant(mutant))

}

func TestProteinToSequence(t *testing.T) {
	assert.Equal(t, "AAAA", proteinToSequence("A"))
}

func TestProteinToSequence2(t *testing.T) {
	assert.Equal(t, "CCCC", proteinToSequence("C"))
}

func TestDnaVerticalToHorizontal(t *testing.T) {
	assert.Equal(t, "AAAAAA", dnaVerticalToHorizontal([]string{"ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF"})[0])
}
func TestDnaVerticalToHorizontal2(t *testing.T) {
	assert.Equal(t, "FFFFFF", dnaVerticalToHorizontal([]string{"ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF", "ABCDEF"})[5])
}

func TestDnaObliqueToHorizontal(t *testing.T) {
	assert.Equal(t, "ATGC", dnaObliqueToHorizontal([]string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}, 0)[0])
}

func TestDnaObliqueToHorizontal2(t *testing.T) {
	assert.Equal(t, "AGTGT", dnaObliqueToHorizontal([]string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}, 1)[3])
}

func TestGenerateArrayOfStrings(t *testing.T) {
	assert.Equal(t, generateArrayOfStrings(2), []string{"", ""})
}

func TestGenerateArrayOfStrings2(t *testing.T) {
	assert.Equal(t, generateArrayOfStrings(6), []string{"", "", "", "", "", ""})
}

func TestDiagonal(t *testing.T) {
	matrix := [adnLength][]string{
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
	}
	assert.Equal(t, diagonal(matrix, 0), "0")
}

func TestDiagonal2(t *testing.T) {
	matrix := [adnLength][]string{
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
	}
	assert.Equal(t, diagonal(matrix, 4), "01234")
}

func TestDiagonal3(t *testing.T) {
	matrix := [adnLength][]string{
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
	}
	assert.Equal(t, diagonal(matrix, 5), "012345")
}

func TestDiagonal4(t *testing.T) {
	matrix := [adnLength][]string{
		{"0", "1", "2", "a", "4", "e"},
		{"0", "1", "2", "b", "d", "5"},
		{"0", "1", "2", "c", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "8", "5"},
		{"0", "1", "2", "8", "4", "5"},
	}
	assert.Equal(t, diagonal(matrix, 8), "885")
}

func TestDiagonal5(t *testing.T) {
	matrix := [adnLength][]string{
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
	}
	assert.Equal(t, diagonal(matrix, 10), "5")
}

func TestDiagonalInverted(t *testing.T) {
	matrix := [adnLength][]string{
		{"a", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"3", "1", "2", "3", "4", "5"},
		{"2", "1", "2", "3", "4", "5"},
		{"A", "1", "2", "3", "4", "5"},
	}
	assert.Equal(t, diagonalInverted(matrix, 0), "A")
}

func TestDiagonalInverted2(t *testing.T) {
	matrix := [adnLength][]string{
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
	}
	assert.Equal(t, diagonalInverted(matrix, 1), "10")
}

func TestDiagonalInverted3(t *testing.T) {
	matrix := [adnLength][]string{
		{"a", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "A"},
	}
	assert.Equal(t, diagonalInverted(matrix, 5), "A4321a")
}

func TestDiagonalInverted6(t *testing.T) {
	matrix := [adnLength][]string{
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "A"},
		{"0", "1", "2", "3", "4", "5"},
	}
	assert.Equal(t, diagonalInverted(matrix, 6), "A4321")
}

func TestDiagonalInverted7(t *testing.T) {
	matrix := [adnLength][]string{
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "5"},
		{"0", "1", "2", "3", "4", "A"},
		{"0", "1", "2", "3", "4", "5"},
	}
	assert.Equal(t, diagonalInverted(matrix, 10), "5")
}
