package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetNetworkInterfacesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetNetworkInterfacesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetNetworkInterfacesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetNetworkInterfacesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetNetworkInterfacesDataSourceMetadata(t *testing.T) {
	d := NewGetNetworkInterfacesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_network_interfaces" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_network_interfaces")
	}
}
