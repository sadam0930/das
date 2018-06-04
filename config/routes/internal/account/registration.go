package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/account"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAccountRegistrationEndpoint = "/api/account/register"
const apiAccountAuthenticationEndpoint = "/api/account/authenticate"

var accountServer = account.AccountServer{
	database.AccountRepository,
	database.OrganizerProvisionRepository,
	database.OrganizerProvisionHistoryRepository,
}

var accountRegistrationController = util.DasController{
	Name:         "AccountRegistrationController",
	Description:  "Create an account in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiAccountRegistrationEndpoint,
	Handler:      accountServer.RegisterAccountHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var accountAuthenticationController = util.DasController{
	Name:         "AccountAuthenticationController",
	Description:  "Authenticate user account",
	Method:       http.MethodPost,
	Endpoint:     apiAccountAuthenticationEndpoint,
	Handler:      accountServer.AccountAuthenticationHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var AccountControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		accountRegistrationController,
		accountAuthenticationController,
	},
}
