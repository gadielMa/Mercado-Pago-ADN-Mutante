package controller

import (
	"net/http"

	"github.com/gadielMa/test/db"
	"github.com/gin-gonic/gin"
)

// json enviado en el GET de /stats
type response struct {
	CountMutantDna int     `json:"count_mutant_dna"`
	CountHumanDna  int     `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}

// Stats godoc
// @Description devuelve la información general de todos los humanos cargados
// @Accept json
// @Produce json
// @Success 200
// @Failure 403
// @Failure 404
// @Failure 500
// @Router /stats [get]
func Stats(ctx *gin.Context) {
	var response response
	response.CountHumanDna = db.GetHumans()
	response.CountMutantDna = db.GetMutants()
	response.Ratio = float64(response.CountMutantDna) / float64(response.CountHumanDna)

	ctx.JSON(http.StatusOK, response)
	return
}
