package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMapGroupDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetMapGroupDataSourceSchemaValidation(t *testing.T) {
	d := NewGetMapGroupDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetMapGroupDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetMapGroupDataSourceMetadata(t *testing.T) {
	d := NewGetMapGroupDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_map_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_map_group")
	}
}
