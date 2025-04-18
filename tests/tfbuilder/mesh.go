package tfbuilder

import (
	"bytes"
	"fmt"
	"text/template"
)

type MeshBuilder struct {
	ResourceName string
	MeshName     string
	Spec         string
	CPID         string // Optional
}

func NewMeshBuilder(resourceName, meshName string) *MeshBuilder {
	return &MeshBuilder{
		ResourceName: resourceName,
		MeshName:     meshName,
	}
}

func (m *MeshBuilder) WithSpec(spec string) *MeshBuilder {
	m.Spec = spec
	return m
}

func (m *MeshBuilder) WithCPID(cpID string) *MeshBuilder {
	m.CPID = cpID
	return m
}

func (m *MeshBuilder) Render(provider ProviderType) string {
	data := map[string]interface{}{
		"Provider":     provider,
		"ResourceName": m.ResourceName,
		"MeshName":     m.MeshName,
		"Spec":         m.Spec,
		"CPID":         m.CPID,
	}

	tmplBytes, err := templatesFS.ReadFile("templates/mesh.tmpl")
	if err != nil {
		panic(fmt.Errorf("failed to read mesh template: %w", err))
	}

	tmpl, err := template.New("mesh").Parse(string(tmplBytes))
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}
