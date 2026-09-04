package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSourceSchemaValidation(t *testing.T) {
	d := NewListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestListFmStoredClusterConfigBackupFileByTaskGroupIdDataSourceMetadata(t *testing.T) {
	d := NewListFmStoredClusterConfigBackupFileByTaskGroupIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_list_fm_stored_cluster_config_backup_file_by_task_group_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_list_fm_stored_cluster_config_backup_file_by_task_group_id")
	}
}
