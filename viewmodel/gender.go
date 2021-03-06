package viewmodel

import (
	"github.com/yubing24/das/businesslogic/reference"
)

type Gender struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GenderDataModelToViewModel(gender reference.Gender) Gender {
	return Gender{
		ID:   gender.ID,
		Name: gender.Name,
	}
}
