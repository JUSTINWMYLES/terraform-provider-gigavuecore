package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllClusterMapsOfAnUserFabricMapListResource)(nil)

// GetAllClusterMapsOfAnUserFabricMapListResource is the generated Terraform list resource implementation.
type GetAllClusterMapsOfAnUserFabricMapListResource struct {
}

// NewGetAllClusterMapsOfAnUserFabricMapListResource returns a new instance of the generated list resource.
func NewGetAllClusterMapsOfAnUserFabricMapListResource() list.ListResource {
	return &GetAllClusterMapsOfAnUserFabricMapListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllClusterMapsOfAnUserFabricMapListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_cluster_maps_of_an_user_fabric_map"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllClusterMapsOfAnUserFabricMapListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get all cluster maps of a user-defined fabric map"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllClusterMapsOfAnUserFabricMapListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
