package provider

import "context"
import (
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	list "github.com/hashicorp/terraform-plugin-framework/list"
	listschema "github.com/hashicorp/terraform-plugin-framework/list/schema"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
)

// Compile-time interface assertion.
var _ list.ListResource = (*ListFmStoredClustersConfigBackupSnapshotsListResource)(nil)

// ListFmStoredClustersConfigBackupSnapshotsListResource is the generated Terraform list resource implementation.
type ListFmStoredClustersConfigBackupSnapshotsListResource struct {
}

// NewListFmStoredClustersConfigBackupSnapshotsListResource returns a new instance of the generated list resource.
func NewListFmStoredClustersConfigBackupSnapshotsListResource() list.ListResource {
	return &ListFmStoredClustersConfigBackupSnapshotsListResource{}
}

// Metadata returns the list resource type name.
func (l *ListFmStoredClustersConfigBackupSnapshotsListResource) Metadata(_ context.Context, _ resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "gigavuecore_list_fm_stored_clusters_config_backup_snapshots"
}

// ListResourceConfigSchema returns the list resource config schema.
func (l *ListFmStoredClustersConfigBackupSnapshotsListResource) ListResourceConfigSchema(_ context.Context, _ list.ListResourceSchemaRequest, resp *list.ListResourceSchemaResponse) {
	resp.Schema = listschema.Schema{MarkdownDescription: "Lists config backup snapshots for multiple clusters"}
}

// List streams matching resource instances for terraform query.
func (l *ListFmStoredClustersConfigBackupSnapshotsListResource) List(_ context.Context, _ list.ListRequest, stream *list.ListResultsStream) {
	stream.Results = list.ListResultsStreamDiagnostics(diag.Diagnostics{diag.NewErrorDiagnostic("Generated provider scaffold", "List is not wired to a remote API endpoint.")})
}
