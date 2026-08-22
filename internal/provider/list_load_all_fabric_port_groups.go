package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllFabricPortGroupsListResource)(nil)

// LoadAllFabricPortGroupsListResource is the generated Terraform list resource implementation.
type LoadAllFabricPortGroupsListResource struct {
}

// NewLoadAllFabricPortGroupsListResource returns a new instance of the generated list resource.
func NewLoadAllFabricPortGroupsListResource() list.ListResource {
	return &LoadAllFabricPortGroupsListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllFabricPortGroupsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_fabric_port_groups"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllFabricPortGroupsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get all the Fabric Port Groups"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllFabricPortGroupsListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
