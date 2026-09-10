package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListBulkReplicateConfigRestoreLogDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestListBulkReplicateConfigRestoreLogDataSourceSchemaValidation(t *testing.T) {
	d := NewListBulkReplicateConfigRestoreLogDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestListBulkReplicateConfigRestoreLogDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestListBulkReplicateConfigRestoreLogDataSourceMetadata(t *testing.T) {
	d := NewListBulkReplicateConfigRestoreLogDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_list_bulk_replicate_config_restore_log" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_list_bulk_replicate_config_restore_log")
	}
}
