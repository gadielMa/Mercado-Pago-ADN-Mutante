package models

// Mutant json recibido en el POST de /mutant
type Mutant struct {
	Dna []string `json:"dna" validate:"eq=6,dive,validate_proteins,min=6,max=6"`
}

// MongoDb estrucutra de la db
type MongoDb struct {
	dna    []string
	mutant bool
}
