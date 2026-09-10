package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllNameServersGroupedByInterfaceDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllNameServersGroupedByInterfaceDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllNameServersGroupedByInterfaceDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllNameServersGroupedByInterfaceDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllNameServersGroupedByInterfaceDataSourceMetadata(t *testing.T) {
	d := NewLoadAllNameServersGroupedByInterfaceDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_name_servers_grouped_by_interface" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_name_servers_grouped_by_interface")
	}
}
