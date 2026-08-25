package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllUsersListResource)(nil)

// LoadAllUsersListResource is the generated Terraform list resource implementation.
type LoadAllUsersListResource struct {
}

// NewLoadAllUsersListResource returns a new instance of the generated list resource.
func NewLoadAllUsersListResource() list.ListResource {
	return &LoadAllUsersListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllUsersListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_users"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllUsersListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load all Users"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllUsersListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
