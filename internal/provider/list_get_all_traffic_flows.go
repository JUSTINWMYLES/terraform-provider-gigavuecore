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
	resp.Schema = listschema.Schema{MarkdownDescription: "List all Traffic Flows with optional filters"}
}

// List streams matching resource instances for terraform query.
func (l *GetAllTrafficFlowsListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
