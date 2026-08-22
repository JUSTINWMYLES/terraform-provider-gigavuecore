package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*DownloadListListResource)(nil)

// DownloadListListResource is the generated Terraform list resource implementation.
type DownloadListListResource struct {
}

// NewDownloadListListResource returns a new instance of the generated list resource.
func NewDownloadListListResource() list.ListResource {
	return &DownloadListListResource{}
}

// Metadata returns the list resource type name.
func (l *DownloadListListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_download_list"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *DownloadListListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Download a nodecryptlist or decryptlist"}
}

// List streams matching resource instances for terraform query.
func (l *DownloadListListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
