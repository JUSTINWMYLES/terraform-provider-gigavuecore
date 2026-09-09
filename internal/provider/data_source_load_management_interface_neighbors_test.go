package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadManagementInterfaceNeighborsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadManagementInterfaceNeighborsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadManagementInterfaceNeighborsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadManagementInterfaceNeighborsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadManagementInterfaceNeighborsDataSourceMetadata(t *testing.T) {
	d := NewLoadManagementInterfaceNeighborsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_management_interface_neighbors" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_management_interface_neighbors")
	}
}
