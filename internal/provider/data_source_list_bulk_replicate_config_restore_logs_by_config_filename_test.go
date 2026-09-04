package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSourceSchemaValidation(t *testing.T) {
	d := NewListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestListBulkReplicateConfigRestoreLogsByConfigFilenameDataSourceMetadata(t *testing.T) {
	d := NewListBulkReplicateConfigRestoreLogsByConfigFilenameDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_list_bulk_replicate_config_restore_logs_by_config_filename" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_list_bulk_replicate_config_restore_logs_by_config_filename")
	}
}
