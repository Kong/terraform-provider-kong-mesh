package tfbuilder

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type ProviderType string

const (
	KongMesh    ProviderType = "kong-mesh"
	Konnect     ProviderType = "konnect"
	KonnectBeta ProviderType = "konnect-beta"
)

type Builder struct {
	provider      ProviderType
	scheme        string
	host          string
	port          int
	region        string
	controlPlanes map[string]*ControlPlane
	meshes        map[string]*MeshBuilder
	policies      map[string]*PolicyBuilder
}

func NewBuilder(provider ProviderType, scheme, host string, port int, region string) *Builder {
	return &Builder{
		provider:      provider,
		scheme:        scheme,
		host:          host,
		port:          port,
		region:        region,
		controlPlanes: make(map[string]*ControlPlane),
		meshes:        make(map[string]*MeshBuilder),
		policies:      make(map[string]*PolicyBuilder),
	}
}

func (b *Builder) AddControlPlane(name string, cp *ControlPlane) *Builder {
	b.controlPlanes[name] = cp
	return b
}

func (b *Builder) AddMesh(mesh *MeshBuilder) *Builder {
	b.meshes[mesh.ResourceName] = mesh
	return b
}

func (b *Builder) RemoveMesh(name string) *Builder {
	delete(b.meshes, name)
	return b
}

func (b *Builder) AddPolicy(p *PolicyBuilder) *Builder {
	b.policies[p.ResourceName] = p
	return b
}

func (b *Builder) RemovePolicy(name string) *Builder {
	delete(b.policies, name)
	return b
}

func (b *Builder) Build() string {
	var sb strings.Builder

	sb.WriteString(b.renderTemplate("provider.tmpl", map[string]interface{}{
		"Provider": b.provider,
		"Scheme":   b.scheme,
		"Host":     b.host,
		"Port":     b.port,
		"Region":   b.region,
	}))
	sb.WriteString("\n")

	for _, cp := range b.controlPlanes {
		sb.WriteString(cp.Render(b.provider))
		sb.WriteString("\n")
	}

	for _, mesh := range b.meshes {
		sb.WriteString(mesh.Render(b.provider))
		sb.WriteString("\n")
	}

	for _, policy := range b.policies {
		sb.WriteString(policy.Render(b.provider))
		sb.WriteString("\n")
	}

	return sb.String()
}

func (b *Builder) renderTemplate(file string, data interface{}) string {
	tmplBytes, err := templatesFS.ReadFile("templates/" + file)
	if err != nil {
		panic(fmt.Errorf("failed to read template %s: %w", file, err))
	}

	tmpl, err := template.New(file).Parse(string(tmplBytes))
	if err != nil {
		panic(fmt.Errorf("failed to parse template: %w", err))
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, data); err != nil {
		panic(fmt.Errorf("failed to render template: %w", err))
	}
	return out.String()
}
