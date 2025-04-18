package tfbuilder

import (
	"bytes"
	"fmt"
	"text/template"
)

type ControlPlane struct {
	ResourceName string
	Name         string
	Description  string
}

func NewControlPlane(resourceName, name, description string) *ControlPlane {
	return &ControlPlane{
		ResourceName: resourceName,
		Name:         name,
		Description:  description,
	}
}

func (cp *ControlPlane) Render(provider ProviderType) string {
	data := map[string]interface{}{
		"Provider":     provider,
		"ResourceName": cp.ResourceName,
		"Name":         cp.Name,
		"Description":  cp.Description,
	}

	tmplBytes, err := templatesFS.ReadFile("templates/control_plane.tmpl")
	if err != nil {
		panic(fmt.Errorf("failed to read control plane template: %w", err))
	}

	tmpl, err := template.New("control_plane").Parse(string(tmplBytes))
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}
