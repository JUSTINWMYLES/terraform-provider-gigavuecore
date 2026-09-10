package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestListBulkReplicateConfigFilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestListBulkReplicateConfigFilesDataSourceSchemaValidation(t *testing.T) {
	d := NewListBulkReplicateConfigFilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestListBulkReplicateConfigFilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestListBulkReplicateConfigFilesDataSourceMetadata(t *testing.T) {
	d := NewListBulkReplicateConfigFilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_list_bulk_replicate_config_files" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_list_bulk_replicate_config_files")
	}
}
