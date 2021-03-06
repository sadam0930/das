package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
)

type ProficiencyServer struct {
	reference.IProficiencyRepository
}

// GET /api/reference/proficiency
func (server ProficiencyServer) SearchProficiencyHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(reference.SearchProficiencyCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", parseErr.Error())
		return
	}

	if proficiencies, err := server.IProficiencyRepository.SearchProficiency(criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	} else {
		dtos := make([]viewmodel.Proficiency, 0)
		for _, each := range proficiencies {
			dtos = append(dtos, viewmodel.ProficiencyDataModelToViewModel(each))
		}
		output, _ := json.Marshal(dtos)
		w.Write(output)
	}
}

// POST /api/reference/proficiency
func (server ProficiencyServer) CreateProficiencyHandler(w http.ResponseWriter, r *http.Request) {}

// DELETE /api/reference/proficiency
func (server ProficiencyServer) DeleteProficiencyHandler(w http.ResponseWriter, r *http.Request) {}

// PUT /api/reference/proficiency
func (server ProficiencyServer) UpdateProficiencyHandler(w http.ResponseWriter, r *http.Request) {}
