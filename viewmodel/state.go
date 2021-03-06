package viewmodel

import "github.com/yubing24/das/businesslogic/reference"

type State struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	CountryID    int    `json:"country"`
}

func StateDataModelToViewModel(dm reference.State) State {
	return State{
		ID:           dm.ID,
		Name:         dm.Name,
		Abbreviation: dm.Abbreviation,
		CountryID:    dm.CountryID,
	}
}
