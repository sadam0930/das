package reference

import (
	"github.com/yubing24/das/businesslogic/reference"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

type CityServer struct {
	reference.ICityRepository
}

// POST /api/reference/city
func (server CityServer) CreateCityHandler(w http.ResponseWriter, r *http.Request) {
	dto := new(viewmodel.CreateCity)
	if err := util.ParseRequestBodyData(r, dto); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	city := dto.ToCityDataModel()

	if err := server.ICityRepository.CreateCity(&city); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}

// DELETE /api/reference/city
func (server CityServer) DeleteCityHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	deleteDTO := new(viewmodel.DeleteCity)
	err := util.ParseRequestBodyData(r, deleteDTO)
	if err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if cities, searchErr := server.ICityRepository.SearchCity(&reference.SearchCityCriteria{CityID: deleteDTO.ID}); searchErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, searchErr.Error(), nil)
		return
	} else if len(cities) != 1 {
		util.RespondJsonResult(w, http.StatusNotFound, "cannot find specified city", nil)
		return
	} else {
		if deleteErr := server.ICityRepository.DeleteCity(cities[0]); deleteErr != nil {
			log.Errorf(ctx, "error in deleting city {ID: %d Name: %s}: %v", cities[0].CityID, cities[0].Name, deleteErr.Error())
			util.RespondJsonResult(w, http.StatusInternalServerError, "cannot delete specified city", nil)
			return
		}
		util.RespondJsonResult(w, http.StatusOK, "success", nil)
		return
	}
}

// PUT /api/reference/city
func (server CityServer) UpdateCityHandler(w http.ResponseWriter, r *http.Request) {
	updateDTO := new(viewmodel.UpdateCity)
	err := util.ParseRequestBodyData(r, updateDTO)
	if err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusNotImplemented, "", nil)
}

// GET /api/reference/city
func (server CityServer) SearchCityHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(reference.SearchCityCriteria)
	err := util.ParseRequestData(r, criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, err.Error())
		return
	}
	cities, err := server.ICityRepository.SearchCity(criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	dtos := make([]viewmodel.City, 0)
	for _, each := range cities {
		dtos = append(dtos, viewmodel.City{
			CityID: each.CityID,
			Name:   each.Name,
			State:  each.StateID,
		})
	}
	output, _ := json.Marshal(dtos)
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
