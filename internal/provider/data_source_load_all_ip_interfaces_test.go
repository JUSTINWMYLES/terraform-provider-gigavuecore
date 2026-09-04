package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllIpInterfacesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllIpInterfacesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllIpInterfacesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllIpInterfacesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllIpInterfacesDataSourceMetadata(t *testing.T) {
	d := NewLoadAllIpInterfacesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_ip_interfaces" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_ip_interfaces")
	}
}
