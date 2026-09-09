package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllInternalFabricMapsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllInternalFabricMapsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllInternalFabricMapsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllInternalFabricMapsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllInternalFabricMapsDataSourceMetadata(t *testing.T) {
	d := NewGetAllInternalFabricMapsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_internal_fabric_maps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_internal_fabric_maps")
	}
}
