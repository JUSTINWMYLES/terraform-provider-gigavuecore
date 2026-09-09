package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFabricPathDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadFabricPathDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadFabricPathDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadFabricPathDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadFabricPathDataSourceMetadata(t *testing.T) {
	d := NewLoadFabricPathDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_fabric_path" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_fabric_path")
	}
}
