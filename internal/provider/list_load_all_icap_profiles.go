package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllIcapProfilesListResource)(nil)

// LoadAllIcapProfilesListResource is the generated Terraform list resource implementation.
type LoadAllIcapProfilesListResource struct {
}

// NewLoadAllIcapProfilesListResource returns a new instance of the generated list resource.
func NewLoadAllIcapProfilesListResource() list.ListResource {
	return &LoadAllIcapProfilesListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllIcapProfilesListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_icap_profiles"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllIcapProfilesListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{Attributes: map[string]listschema.Attribute{"cluster_id": listschema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "page": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "sort": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllIcapProfilesListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
