package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllSffpProfileListResource)(nil)

// GetAllSffpProfileListResource is the generated Terraform list resource implementation.
type GetAllSffpProfileListResource struct {
}

// NewGetAllSffpProfileListResource returns a new instance of the generated list resource.
func NewGetAllSffpProfileListResource() list.ListResource {
	return &GetAllSffpProfileListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllSffpProfileListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_sffp_profile"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllSffpProfileListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "new in H5.8"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllSffpProfileListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
