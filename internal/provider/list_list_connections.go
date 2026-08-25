package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*ListConnectionsListResource)(nil)

// ListConnectionsListResource is the generated Terraform list resource implementation.
type ListConnectionsListResource struct {
}

// NewListConnectionsListResource returns a new instance of the generated list resource.
func NewListConnectionsListResource() list.ListResource {
	return &ListConnectionsListResource{}
}

// Metadata returns the list resource type name.
func (l *ListConnectionsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_list_connections"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *ListConnectionsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "List all unified resource connections by environment id"}
}

// List streams matching resource instances for terraform query.
func (l *ListConnectionsListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
