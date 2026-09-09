package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListFmStoredClustersConfigBackupSnapshotsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestListFmStoredClustersConfigBackupSnapshotsDataSourceSchemaValidation(t *testing.T) {
	d := NewListFmStoredClustersConfigBackupSnapshotsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestListFmStoredClustersConfigBackupSnapshotsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestListFmStoredClustersConfigBackupSnapshotsDataSourceMetadata(t *testing.T) {
	d := NewListFmStoredClustersConfigBackupSnapshotsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_list_fm_stored_clusters_config_backup_snapshots" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_list_fm_stored_clusters_config_backup_snapshots")
	}
}
