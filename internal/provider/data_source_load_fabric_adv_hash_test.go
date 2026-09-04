package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFabricAdvHashDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadFabricAdvHashDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadFabricAdvHashDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadFabricAdvHashDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadFabricAdvHashDataSourceMetadata(t *testing.T) {
	d := NewLoadFabricAdvHashDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_fabric_adv_hash" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_fabric_adv_hash")
	}
}
