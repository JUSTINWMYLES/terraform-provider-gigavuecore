package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentPosIdsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCurrentPosIdsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCurrentPosIdsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCurrentPosIdsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCurrentPosIdsDataSourceMetadata(t *testing.T) {
	d := NewGetCurrentPosIdsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_current_pos_ids" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_current_pos_ids")
	}
}
