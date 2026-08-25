package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllMetadataNetworkProfilesListResource)(nil)

// LoadAllMetadataNetworkProfilesListResource is the generated Terraform list resource implementation.
type LoadAllMetadataNetworkProfilesListResource struct {
}

// NewLoadAllMetadataNetworkProfilesListResource returns a new instance of the generated list resource.
func NewLoadAllMetadataNetworkProfilesListResource() list.ListResource {
	return &LoadAllMetadataNetworkProfilesListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllMetadataNetworkProfilesListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_metadata_network_profiles"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllMetadataNetworkProfilesListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load All Metadata Network Profiles"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllMetadataNetworkProfilesListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
