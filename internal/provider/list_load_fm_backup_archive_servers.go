package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*LoadFmBackupArchiveServersListResource)(nil)

// LoadFmBackupArchiveServersListResource is the generated Terraform list resource implementation.
type LoadFmBackupArchiveServersListResource struct {
}

// NewLoadFmBackupArchiveServersListResource returns a new instance of the generated list resource.
func NewLoadFmBackupArchiveServersListResource() list.ListResource {
	return &LoadFmBackupArchiveServersListResource{}
}

// Metadata returns the list resource type name.
func (l *LoadFmBackupArchiveServersListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_load_fm_backup_archive_servers"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *LoadFmBackupArchiveServersListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "List FM backup archive servers information"}
}

// List streams matching resource instances for terraform query.
func (l *LoadFmBackupArchiveServersListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
