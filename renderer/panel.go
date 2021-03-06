package renderer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/josephbudd/kickwasm/cases"
	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

// createGoPanels creates the renderer/panels/ go panel files.
// Only for organic not autogenerated panels.
func createGoPanels(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	folderpaths := appPaths.GetPaths()
	servicePanelNamePathMap := builder.GenerateServiceEmptyInsidePanelNamePathMap()
	serviceButtonPanelGroups := builder.GenerateServiceButtonPanelGroups()
	serviceNames := builder.GenerateOrganicServiceNames()
	for _, serviceName := range serviceNames {
		panelNamePathMap := servicePanelNamePathMap[serviceName]
		serviceButtonPanelGroup := serviceButtonPanelGroups[serviceName]
		for _, buttonPanelGroups := range serviceButtonPanelGroup {
			// make this panel's group
			panelGroup := make([]*tap.Panel, 0, 5)
			for _, panel := range buttonPanelGroups.PanelNamesIDMap {
				panelGroup = append(panelGroup, panel)
			}
			// template data for each panel file in this group.
			for panelName, panel := range buttonPanelGroups.PanelNamesIDMap {
				folders := strings.Join(panelNamePathMap[panelName], string(os.PathSeparator))
				folderpath := filepath.Join(folderpaths.OutputRendererPanels, folders, panelName)
				if err := os.MkdirAll(folderpath, appPaths.GetDMode()); err != nil {
					return err
				}
				data := &struct {
					PanelName                          string
					PanelID                            string
					PanelGroup                         []*tap.Panel
					IsTabSiblingPanel                  bool
					ApplicationGitPath                 string
					ImportRendererViewTools            string
					ImportDomainTypes                  string
					ImportDomainImplementationsCalling string

					CamelCase      func(string) string
					LowerCamelCase func(string) string
					SplitTabJoin   func(string) string
				}{
					PanelName:                          panelName,
					PanelID:                            panel.HTMLID,
					PanelGroup:                         panelGroup,
					IsTabSiblingPanel:                  buttonPanelGroups.IsTabButton,
					ApplicationGitPath:                 builder.ImportPath,
					ImportRendererViewTools:            folderpaths.ImportRendererViewTools,
					ImportDomainTypes:                  folderpaths.ImportDomainTypes,
					ImportDomainImplementationsCalling: folderpaths.ImportDomainImplementationsCalling,

					CamelCase:      cases.CamelCase,
					LowerCamelCase: cases.LowerCamelCase,
					SplitTabJoin: func(s string) string {
						ss := strings.Split(s, "\n")
						return "\t" + strings.Join(ss, "\n\t")
					},
				}
				fname := "panel.go"
				oPath := filepath.Join(folderpath, fname)
				if err := templates.ProcessTemplate(fname, oPath, templates.Panel, data, appPaths); err != nil {
					return err
				}
				fname = "controler.go"
				oPath = filepath.Join(folderpath, fname)
				if err := templates.ProcessTemplate(fname, oPath, templates.PanelControler, data, appPaths); err != nil {
					return err
				}
				fname = "presenter.go"
				oPath = filepath.Join(folderpath, fname)
				if err := templates.ProcessTemplate(fname, oPath, templates.PanelPresenter, data, appPaths); err != nil {
					return err
				}
				fname = "caller.go"
				oPath = filepath.Join(folderpath, fname)
				if err := templates.ProcessTemplate(fname, oPath, templates.PanelCaller, data, appPaths); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
