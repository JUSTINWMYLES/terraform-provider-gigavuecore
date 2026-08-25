package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllV3UsersListResource)(nil)

// LoadAllV3UsersListResource is the generated Terraform list resource implementation.
type LoadAllV3UsersListResource struct {
}

// NewLoadAllV3UsersListResource returns a new instance of the generated list resource.
func NewLoadAllV3UsersListResource() list.ListResource {
	return &LoadAllV3UsersListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllV3UsersListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_v3_users"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllV3UsersListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load all SNMPv3 Users"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllV3UsersListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
