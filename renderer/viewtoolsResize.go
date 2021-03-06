package renderer

import (
	"path/filepath"

	"github.com/josephbudd/kickwasm/paths"
	"github.com/josephbudd/kickwasm/renderer/templates"
	"github.com/josephbudd/kickwasm/tap"
)

func createResizeGo(appPaths paths.ApplicationPathsI, builder *tap.Builder) error {
	data := &struct {
		IDs        *tap.IDs
		Classes    *tap.Classes
		Attributes *tap.Attributes
	}{
		IDs:        builder.IDs,
		Classes:    builder.Classes,
		Attributes: builder.Attributes,
	}
	// execute the template
	folderpaths := appPaths.GetPaths()
	fname := "resize.go"
	oPath := filepath.Join(folderpaths.OutputRendererViewTools, fname)
	return templates.ProcessTemplate(fname, oPath, templates.ViewToolsResize, data, appPaths)
}
