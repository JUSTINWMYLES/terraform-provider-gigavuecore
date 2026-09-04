package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFabricPortGroupsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllFabricPortGroupsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllFabricPortGroupsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllFabricPortGroupsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllFabricPortGroupsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllFabricPortGroupsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_fabric_port_groups" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_fabric_port_groups")
	}
}
