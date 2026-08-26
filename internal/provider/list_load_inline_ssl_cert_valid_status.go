package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadInlineSslCertValidStatusListResource)(nil)

// LoadInlineSslCertValidStatusListResource is the generated Terraform list resource implementation.
type LoadInlineSslCertValidStatusListResource struct {
}

// NewLoadInlineSslCertValidStatusListResource returns a new instance of the generated list resource.
func NewLoadInlineSslCertValidStatusListResource() list.ListResource {
	return &LoadInlineSslCertValidStatusListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadInlineSslCertValidStatusListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_inline_ssl_cert_valid_status"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadInlineSslCertValidStatusListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get Inline SSL certificate validation cache status"}
}

// List streams matching resource instances for terraform query.
func (l *LoadInlineSslCertValidStatusListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
