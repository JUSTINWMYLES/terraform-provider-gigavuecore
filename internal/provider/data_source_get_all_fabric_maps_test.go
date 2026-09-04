package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFabricMapsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllFabricMapsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllFabricMapsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllFabricMapsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllFabricMapsDataSourceMetadata(t *testing.T) {
	d := NewGetAllFabricMapsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_fabric_maps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_fabric_maps")
	}
}
