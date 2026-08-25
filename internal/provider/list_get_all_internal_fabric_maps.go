package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllInternalFabricMapsListResource)(nil)

// GetAllInternalFabricMapsListResource is the generated Terraform list resource implementation.
type GetAllInternalFabricMapsListResource struct {
}

// NewGetAllInternalFabricMapsListResource returns a new instance of the generated list resource.
func NewGetAllInternalFabricMapsListResource() list.ListResource {
	return &GetAllInternalFabricMapsListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllInternalFabricMapsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_internal_fabric_maps"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllInternalFabricMapsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get all internally generated fabric maps supporting a specific user-defined fabric map"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllInternalFabricMapsListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
