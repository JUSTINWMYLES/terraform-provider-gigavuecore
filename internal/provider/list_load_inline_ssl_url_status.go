package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadInlineSslUrlStatusListResource)(nil)

// LoadInlineSslUrlStatusListResource is the generated Terraform list resource implementation.
type LoadInlineSslUrlStatusListResource struct {
}

// NewLoadInlineSslUrlStatusListResource returns a new instance of the generated list resource.
func NewLoadInlineSslUrlStatusListResource() list.ListResource {
	return &LoadInlineSslUrlStatusListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadInlineSslUrlStatusListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_inline_ssl_url_status"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadInlineSslUrlStatusListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get Inline SSL URL cache status"}
}

// List streams matching resource instances for terraform query.
func (l *LoadInlineSslUrlStatusListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
