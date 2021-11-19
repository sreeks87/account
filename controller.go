package account

import (
	"net/http"
)

// The controller strcut which has an object of the service
type accountController struct {
	svc Service
}

// NewAcctController returns the object of controller
// this is what the user uses to interact with the fucntionalities of the package
func NewAcctController(base string) *accountController {
	return &accountController{
		svc: NewAccountService(http.DefaultClient, base),
	}
}

// CreateAccount function in the controller calls the service to get the account created.
func (a *accountController) CreateAccount(data Data) (Data, error) {
	return a.svc.Create(&data)
}

// DeleteAccount function in the controller calls the service to get the account deleted.
func (a *accountController) DeleteAccount(acid string, version string) error {
	return a.svc.Delete(acid, version)
}

// FetchAccount function in the controller calls the service to get the account details.
func (a *accountController) FetchAccount(acid string) (Data, error) {
	return a.svc.Fetch(acid)
}
