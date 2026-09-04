package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListFmStoredClusterConfigBackupSnapshotDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestListFmStoredClusterConfigBackupSnapshotDataSourceSchemaValidation(t *testing.T) {
	d := NewListFmStoredClusterConfigBackupSnapshotDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestListFmStoredClusterConfigBackupSnapshotDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestListFmStoredClusterConfigBackupSnapshotDataSourceMetadata(t *testing.T) {
	d := NewListFmStoredClusterConfigBackupSnapshotDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_list_fm_stored_cluster_config_backup_snapshot" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_list_fm_stored_cluster_config_backup_snapshot")
	}
}
