package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMapsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllMapsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllMapsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllMapsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllMapsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllMapsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_maps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_maps")
	}
}
