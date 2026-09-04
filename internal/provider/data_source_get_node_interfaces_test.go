package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNodeInterfacesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetNodeInterfacesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetNodeInterfacesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetNodeInterfacesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetNodeInterfacesDataSourceMetadata(t *testing.T) {
	d := NewGetNodeInterfacesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_node_interfaces" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_node_interfaces")
	}
}
