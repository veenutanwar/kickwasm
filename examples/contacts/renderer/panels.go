package main

import (
	"github.com/josephbudd/kicknotjs"
	"github.com/josephbudd/kickwasm/examples/contacts/domain/types"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/panels/AboutButton/AboutPanel"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/panels/AddButton/AddContactPanel"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/panels/EditButton/EditContactEditPanel"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/panels/EditButton/EditContactNotReadyPanel"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/panels/EditButton/EditContactSelectPanel"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/panels/RemoveButton/RemoveContactConfirmPanel"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/panels/RemoveButton/RemoveContactNotReadyPanel"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/panels/RemoveButton/RemoveContactSelectPanel"
	"github.com/josephbudd/kickwasm/examples/contacts/renderer/viewtools"
)

/*

	DO NOT EDIT THIS FILE.

	This file is generated by kickasm and regenerated by rekickasm.

*/

func doPanels(quitCh chan struct{}, tools *viewtools.Tools, callMap types.RendererCallMap, notjs *kicknotjs.NotJS) {
	// 1. Construct the panel code.
	addContactPanel := AddContactPanel.NewPanel(quitCh, tools, notjs, callMap)
	editContactEditPanel := EditContactEditPanel.NewPanel(quitCh, tools, notjs, callMap)
	editContactNotReadyPanel := EditContactNotReadyPanel.NewPanel(quitCh, tools, notjs, callMap)
	editContactSelectPanel := EditContactSelectPanel.NewPanel(quitCh, tools, notjs, callMap)
	removeContactConfirmPanel := RemoveContactConfirmPanel.NewPanel(quitCh, tools, notjs, callMap)
	removeContactNotReadyPanel := RemoveContactNotReadyPanel.NewPanel(quitCh, tools, notjs, callMap)
	removeContactSelectPanel := RemoveContactSelectPanel.NewPanel(quitCh, tools, notjs, callMap)
	aboutPanel := AboutPanel.NewPanel(quitCh, tools, notjs, callMap)

	// 2. Size the app.
	tools.SizeApp()

	// 3. Start each panel's initial calls.
	addContactPanel.InitialCalls()
	editContactEditPanel.InitialCalls()
	editContactNotReadyPanel.InitialCalls()
	editContactSelectPanel.InitialCalls()
	removeContactConfirmPanel.InitialCalls()
	removeContactNotReadyPanel.InitialCalls()
	removeContactSelectPanel.InitialCalls()
	aboutPanel.InitialCalls()
}