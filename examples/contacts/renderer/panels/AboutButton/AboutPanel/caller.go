package AboutPanel

import (
	"github.com/josephbudd/kickwasm/examples/contacts/domain/implementations/calling"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

// Caller communicates with the main process.
type Caller struct {
	presenter *Presenter
	quitCh    chan struct{}
	connection types.RendererCallMap
	tools     *viewtools.Tools
}

func (panelCaller *Caller) addCallBacks() {
	getAboutCall := panelCaller.connection[calling.GetAboutCallID]
	getAboutCall.AddCallBack(panelCaller.getAboutCB)
}

func (panelCaller *Caller) initialCalls() {
	getAboutCall := panelCaller.connection[calling.GetAboutCallID]
	getAboutCall.CallMainProcess(nil)
}

func (panelCaller *Caller) getAboutCB(params interface{}) {
	switch params := params.(type) {
	case *calling.MainProcessToRendererGetAboutParams:
		if params.Error {
			panelCaller.tools.Error(params.ErrorMessage)
			return
		}
		panelCaller.presenter.displayReleases(params.Version, params.Releases)
		panelCaller.presenter.displayContributors(params.Contributors)
		panelCaller.presenter.displayCredits(params.Credits)
		panelCaller.presenter.displayLicenses(params.Licenses)
	}
}