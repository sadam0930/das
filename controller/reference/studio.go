package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"net/http"
)

type StudioServer struct {
	reference.IStudioRepository
}

// GET /api/reference/studio
func (server StudioServer) SearchStudioHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(reference.SearchStudioCriteria)

	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	if studios, err := server.SearchStudio(criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	} else {
		data := make([]viewmodel.Studio, 0)
		for _, each := range studios {
			data = append(data, viewmodel.StudioDataModelToViewModel(each))
		}

		output, _ := json.Marshal(data)
		w.Write(output)
	}
}

// POST /api/reference/studio
func (server StudioServer) CreateStudioHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}

// PUT /api/reference/studio
func (server StudioServer) UpdateStudioHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}

// DELETE /api/reference/studio
func (server StudioServer) DeleteStudioHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}
