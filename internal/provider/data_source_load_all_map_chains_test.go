package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMapChainsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllMapChainsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllMapChainsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllMapChainsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllMapChainsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllMapChainsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_map_chains" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_map_chains")
	}
}
