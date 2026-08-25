package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadAllSslDecryptionKeyMapsListResource)(nil)

// LoadAllSslDecryptionKeyMapsListResource is the generated Terraform list resource implementation.
type LoadAllSslDecryptionKeyMapsListResource struct {
}

// NewLoadAllSslDecryptionKeyMapsListResource returns a new instance of the generated list resource.
func NewLoadAllSslDecryptionKeyMapsListResource() list.ListResource {
	return &LoadAllSslDecryptionKeyMapsListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadAllSslDecryptionKeyMapsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_all_ssl_decryption_key_maps"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadAllSslDecryptionKeyMapsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Load all defined SSL Decryption KeyMaps"}
}

// List streams matching resource instances for terraform query.
func (l *LoadAllSslDecryptionKeyMapsListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
