package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllTacacsServersListResource)(nil)

// LoadAllTacacsServersListResource is the generated Terraform list resource implementation.
type LoadAllTacacsServersListResource struct {
}

// NewLoadAllTacacsServersListResource returns a new instance of the generated list resource.
func NewLoadAllTacacsServersListResource() list.ListResource {
	return &LoadAllTacacsServersListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllTacacsServersListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_tacacs_servers"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllTacacsServersListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load all TACACS+ Servers"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllTacacsServersListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
