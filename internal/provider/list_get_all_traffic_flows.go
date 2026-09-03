package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetAllTrafficFlowsListResource)(nil)

// GetAllTrafficFlowsListResource is the generated Terraform list resource implementation.
type GetAllTrafficFlowsListResource struct {
}

// NewGetAllTrafficFlowsListResource returns a new instance of the generated list resource.
func NewGetAllTrafficFlowsListResource() list.ListResource {
	return &GetAllTrafficFlowsListResource{}
}

// Metadata returns the list resource type name.
func (l *GetAllTrafficFlowsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_all_traffic_flows"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetAllTrafficFlowsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "List all Traffic Flows with optional filters", Attributes: map[string]listschema.Attribute{"alias": listschema.StringAttribute{MarkdownDescription: "Traffic Flows alias filter", Optional: true}, "config_status": listschema.StringAttribute{MarkdownDescription: "Configuration status filter", Optional: true}, "dst_cluster": listschema.StringAttribute{MarkdownDescription: "Destination cluster filter", Optional: true}, "dst_ports": listschema.StringAttribute{MarkdownDescription: "Destination port(s) filter", Optional: true}, "health_state": listschema.StringAttribute{MarkdownDescription: "Health state filter", Optional: true}, "page": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed pair of values in a (pageNo:pageSize) format. 'pageNo' is 1-based. If omitted, entire list of entities is returned", Optional: true}, "select": listschema.StringAttribute{MarkdownDescription: "Comma-separated list of fields to select", Optional: true}, "sort": listschema.StringAttribute{MarkdownDescription: "parentheses-enclosed comma-separated list of entity attributes, optionally qualified with the sort order attribute. The default sort order is ASC and default sort field is fabric map alias. Example: sort=(aaa,bbb:ASC,ccc:DESC)", Optional: true}, "src_cluster": listschema.StringAttribute{MarkdownDescription: "Source cluster filter", Optional: true}, "src_ports": listschema.StringAttribute{MarkdownDescription: "Source port(s) filter", Optional: true}, "summary": listschema.BoolAttribute{MarkdownDescription: "Return summary data", Optional: true}, "type": listschema.StringAttribute{MarkdownDescription: "Traffic Flows type filter", Optional: true}}}
}

// List streams matching resource instances for terraform query.
func (l *GetAllTrafficFlowsListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
