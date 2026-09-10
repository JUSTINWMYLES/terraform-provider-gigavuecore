package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadNrtStatsFabricMapDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadNrtStatsFabricMapDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadNrtStatsFabricMapDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadNrtStatsFabricMapDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadNrtStatsFabricMapDataSourceMetadata(t *testing.T) {
	d := NewLoadNrtStatsFabricMapDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_nrt_stats_fabric_map" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_nrt_stats_fabric_map")
	}
}
