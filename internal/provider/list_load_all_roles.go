package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllRolesListResource)(nil)

// LoadAllRolesListResource is the generated Terraform list resource implementation.
type LoadAllRolesListResource struct {
}

// NewLoadAllRolesListResource returns a new instance of the generated list resource.
func NewLoadAllRolesListResource() list.ListResource {
	return &LoadAllRolesListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllRolesListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_roles"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllRolesListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load all roles"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllRolesListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
