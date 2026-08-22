package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllToolPortMirrorListResource)(nil)

// LoadAllToolPortMirrorListResource is the generated Terraform list resource implementation.
type LoadAllToolPortMirrorListResource struct {
}

// NewLoadAllToolPortMirrorListResource returns a new instance of the generated list resource.
func NewLoadAllToolPortMirrorListResource() list.ListResource {
	return &LoadAllToolPortMirrorListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllToolPortMirrorListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_tool_port_mirror"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllToolPortMirrorListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load all port mirrors"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllToolPortMirrorListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
