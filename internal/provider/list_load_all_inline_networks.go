package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllInlineNetworksListResource)(nil)

// LoadAllInlineNetworksListResource is the generated Terraform list resource implementation.
type LoadAllInlineNetworksListResource struct {
}

// NewLoadAllInlineNetworksListResource returns a new instance of the generated list resource.
func NewLoadAllInlineNetworksListResource() list.ListResource {
	return &LoadAllInlineNetworksListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllInlineNetworksListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_inline_networks"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllInlineNetworksListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load all Inline Networks"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllInlineNetworksListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
