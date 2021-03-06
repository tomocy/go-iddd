package presenter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
	derr "github.com/tomocy/archs/domain/error"
	"github.com/tomocy/archs/domain/model"
	"github.com/tomocy/archs/infra/http/route"
	"github.com/tomocy/archs/infra/http/session"
	"github.com/tomocy/archs/infra/http/view"
)

func New(view view.View, w http.ResponseWriter, r *http.Request) *Presenter {
	return &Presenter{
		view:       view,
		respWriter: w,
		request:    r,
	}
}

type Presenter struct {
	view       view.View
	respWriter http.ResponseWriter
	request    *http.Request
}

func (p *Presenter) ShowUserRegistrationForm() {
	if err := p.showWitDefaults("user.new", nil); err != nil {
		p.logInternalServerError("show user registration form", err)
	}
}

func (p *Presenter) OnUserRegistered(user *model.User) {
	dest := fmt.Sprintf("%s/%s", route.Web.Route("user.show"), user.ID)
	log.Printf("register user successfully: %v\n", user)
	p.redirect(dest)
}

func (p *Presenter) OnUserRegistrationFailed(err error) {
	cause := errors.Cause(err)
	switch {
	case derr.InInput(cause):
		session.SetErrorMessage(p.respWriter, p.request, cause.Error())
		p.redirect(route.Web.Route("user.new").String())
	default:
		p.logUnknownError("user registration", cause)
	}
}

func (p *Presenter) OnUserFindingFailed(err error) {
	cause := errors.Cause(err)
	switch {
	case derr.InInput(cause):
		p.respWriter.WriteHeader(http.StatusNotFound)
	default:
		p.logUnknownError("user finding", cause)
	}
}

func (p *Presenter) showWitDefaults(name string, data map[string]interface{}) error {
	if data == nil {
		data = make(map[string]interface{})
	}
	data["Error"] = session.GetErrorMessage(p.respWriter, p.request)

	return p.show(name, data)
}

func (p *Presenter) show(name string, data interface{}) error {
	return p.view.Show(p.respWriter, name, data)
}

func (p *Presenter) redirect(dest string) {
	http.Redirect(p.respWriter, p.request, dest, http.StatusSeeOther)
}

func (p *Presenter) logInternalServerError(did string, msg interface{}) {
	log.Printf("failed to %s: %v\n", did, msg)
	p.respWriter.WriteHeader(http.StatusInternalServerError)
}

func (p *Presenter) logUnknownError(did string, err error) {
	log.Printf("failed to deal with unknown error in %s: %v\n", did, err)
	p.respWriter.WriteHeader(http.StatusInternalServerError)
}
