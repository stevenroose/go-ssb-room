package errors

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/go-kit/kit/log/level"
	"go.mindeco.de/http/auth"
	"go.mindeco.de/http/render"
	"go.mindeco.de/logging"

	"github.com/ssb-ngi-pointer/go-ssb-room/roomdb"
	"github.com/ssb-ngi-pointer/go-ssb-room/web/i18n"
)

type ErrorHandler struct {
	locHelper *i18n.Helper
	render    *render.Renderer
}

func NewErrorHandler(locHelper *i18n.Helper) *ErrorHandler {
	return &ErrorHandler{
		locHelper: locHelper,
	}
}

// SetRenderer needs to update the rendere later since we need to pass ErrorHandler into render.New (ie. befor we get the pointer for r)
func (eh *ErrorHandler) SetRenderer(r *render.Renderer) {
	eh.render = r
}

func (eh *ErrorHandler) Handle(rw http.ResponseWriter, req *http.Request, code int, err error) {
	var ih = i18n.LocalizerFromRequest(eh.locHelper, req)

	// default, unlocalized message
	msg := err.Error()

	// localize some specific error messages
	var (
		aa  roomdb.ErrAlreadyAdded
		pnf PageNotFound
		br  ErrBadRequest
	)

	switch {
	case err == auth.ErrBadLogin:
		msg = ih.LocalizeSimple("ErrorAuthBadLogin")

	case errors.Is(err, roomdb.ErrNotFound):
		msg = ih.LocalizeSimple("ErrorNotFound")

	case errors.As(err, &aa):
		msg = ih.LocalizeWithData("ErrorAlreadyAdded", map[string]string{
			"Feed": aa.Ref.Ref(),
		})

	case errors.As(err, &pnf):
		msg = ih.LocalizeWithData("ErrorPageNotFound", map[string]string{
			"Path": pnf.Path,
		})

	case errors.As(err, &br):
		// TODO: we could localize all the "Where:" as labels, too
		// buttt it feels like overkill right now
		msg = ih.LocalizeWithData("ErrorBadRequest", map[string]string{
			"Where":   br.Where,
			"Details": br.Details.Error(),
		})
	}

	data := errorTemplateData{
		Err: template.HTML(msg),
		// TODO: localize status codes? might be fine with a few
		Status:     http.StatusText(code),
		StatusCode: code,
	}

	renderErr := eh.render.Render(rw, req, "error.tmpl", code, data)
	if renderErr != nil {
		logger := logging.FromContext(req.Context())
		level.Error(logger).Log("event", "error template renderfailed",
			"orig-err", err,
			"render-err", renderErr,
		)
	}
}

type errorTemplateData struct {
	StatusCode int
	Status     string
	Err        template.HTML
}
