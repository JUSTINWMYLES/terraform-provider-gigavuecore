package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllInlineToolGroupsListResource)(nil)

// LoadAllInlineToolGroupsListResource is the generated Terraform list resource implementation.
type LoadAllInlineToolGroupsListResource struct {
}

// NewLoadAllInlineToolGroupsListResource returns a new instance of the generated list resource.
func NewLoadAllInlineToolGroupsListResource() list.ListResource {
	return &LoadAllInlineToolGroupsListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllInlineToolGroupsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_inline_tool_groups"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllInlineToolGroupsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load all Inline Tool Groups"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllInlineToolGroupsListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
