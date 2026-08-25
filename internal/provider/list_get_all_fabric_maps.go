package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllFabricMapsListResource)(nil)

// GetAllFabricMapsListResource is the generated Terraform list resource implementation.
type GetAllFabricMapsListResource struct {
}

// NewGetAllFabricMapsListResource returns a new instance of the generated list resource.
func NewGetAllFabricMapsListResource() list.ListResource {
	return &GetAllFabricMapsListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllFabricMapsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_fabric_maps"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllFabricMapsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get all user-defined fabric maps"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllFabricMapsListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
