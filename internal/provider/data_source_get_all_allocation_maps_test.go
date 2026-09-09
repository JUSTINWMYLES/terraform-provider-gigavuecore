package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAllocationMapsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllAllocationMapsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllAllocationMapsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllAllocationMapsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllAllocationMapsDataSourceMetadata(t *testing.T) {
	d := NewGetAllAllocationMapsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_allocation_maps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_allocation_maps")
	}
}
