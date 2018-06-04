package reference

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiReferenceSchoolEndpoint = "/api/reference/school"

var schoolServer = reference.SchoolServer{
	database.SchoolRepository,
}

var searchSchoolController = util.DasController{
	Name:         "SearchSchoolController",
	Description:  "Search schools in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceSchoolEndpoint,
	Handler:      schoolServer.SearchSchoolHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createSchoolController = util.DasController{
	Name:         "CreateSchoolController",
	Description:  "Create a school in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceSchoolEndpoint,
	Handler:      schoolServer.CreateSchoolHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR, businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var deleteSchoolController = util.DasController{
	Name:         "DeleteSchoolController",
	Description:  "Delete a school from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceSchoolEndpoint,
	Handler:      schoolServer.DeleteSchoolHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateSchoolController = util.DasController{
	Name:         "UpdateSchoolController",
	Description:  "Update a school in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceSchoolEndpoint,
	Handler:      schoolServer.UpdateSchoolHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var SchoolControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchSchoolController,
		createSchoolController,
		deleteSchoolController,
		updateSchoolController,
	},
}
