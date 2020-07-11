package models

// Request json recibido en el POST de /mutant
type Request struct {
	Dna []string `json:"dna" validate:"eq=6,dive,validate_proteins,min=6,max=6"`
}

// Response json enviado en el GET de /stats
type Response struct {
	CountMutantDna int     `json:"count_mutant_dna"`
	CountHumanDna  int     `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}
