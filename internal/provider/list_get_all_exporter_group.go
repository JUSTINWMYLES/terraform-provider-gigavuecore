package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllExporterGroupListResource)(nil)

// GetAllExporterGroupListResource is the generated Terraform list resource implementation.
type GetAllExporterGroupListResource struct {
}

// NewGetAllExporterGroupListResource returns a new instance of the generated list resource.
func NewGetAllExporterGroupListResource() list.ListResource {
	return &GetAllExporterGroupListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllExporterGroupListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_exporter_group"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllExporterGroupListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get all Apps Exporter Group"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllExporterGroupListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
