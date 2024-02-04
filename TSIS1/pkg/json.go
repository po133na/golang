package tsis1

import (
	"errors"
)

type Breeds struct {
	Id        string `json:"Id"`
	BreedName string `json:"Breedname"`
	Country   string `json:"Country"`
	Desc      string `json:"Desc"`
}

var breeds = []Breeds{
	{
		"Id":        "1",
		"BreedName": "Persian",
		"Country":   "Iran",
		"Desc":      "Long-haired cat breed",
	},
	{
		"Id":        "2",
		"BreedName": "Siamese",
		"Country":   "Thailand",
		"Desc":      "Sleek and short-haired cat breed",
	},
	{
		"Id":        "3",
		"BreedName": "Maine Coon",
		"Country":   "United States",
		"Desc":      "One of the largest domesticated cat breeds",
	},
	{
		"Id":        "4",
		"BreedName": "Bengal",
		"Country":   "United States",
		"Desc":      "Distinctive spotted or marbled coat pattern",
	},
	{
		"Id":        "5",
		"BreedName": "Scottish Fold",
		"Country":   "Scotland",
		"Desc":      "Known for its folded ears",
	},
	{
		"Id":        "6",
		"BreedName": "Ragdoll",
		"Country":   "United States",
		"Desc":      "Large, long-haired cat with striking blue eyes",
	},
	{
		"Id":        "7",
		"BreedName": "Sphynx",
		"Country":   "Canada",
		"Desc":      "Hairless cat breed",
	},
	{
		"Id":        "8",
		"BreedName": "Abyssinian",
		"Country":   "Ethiopia",
		"Desc":      "Small to medium-sized, short-haired cat",
	},
	{
		"Id":        "9",
		"BreedName": "British Shorthair",
		"Country":   "United Kingdom",
		"Desc":      "Round face and dense coat",
	},
	{
		"Id":        "10",
		"BreedName": "Norwegian Forest",
		"Country":   "Norway",
		"Desc":      "Large, sturdy cat with a semi-long coat",
	},
}

func GetBreeds() []Breeds {
	return breeds
}

func GetCat(id string, breeds []Breeds) {
	for _, c := range breeds {
		if c.ID == id {
			return &c, nil
		}
	}
	return 0, errors.New("Cat not found")
}
