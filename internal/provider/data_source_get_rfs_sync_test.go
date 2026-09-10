package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetRfsSyncDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetRfsSyncDataSourceSchemaValidation(t *testing.T) {
	d := NewGetRfsSyncDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetRfsSyncDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetRfsSyncDataSourceMetadata(t *testing.T) {
	d := NewGetRfsSyncDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_rfs_sync" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_rfs_sync")
	}
}
