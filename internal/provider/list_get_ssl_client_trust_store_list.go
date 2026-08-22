package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*GetSslClientTrustStoreListListResource)(nil)

// GetSslClientTrustStoreListListResource is the generated Terraform list resource implementation.
type GetSslClientTrustStoreListListResource struct {
}

// NewGetSslClientTrustStoreListListResource returns a new instance of the generated list resource.
func NewGetSslClientTrustStoreListListResource() list.ListResource {
	return &GetSslClientTrustStoreListListResource{}
}

// Metadata returns the list resource type name.
func (l *GetSslClientTrustStoreListListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_get_ssl_client_trust_store_list"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *GetSslClientTrustStoreListListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Get information about the SSL Client trust-stores"}
}

// List streams matching resource instances for terraform query.
func (l *GetSslClientTrustStoreListListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
