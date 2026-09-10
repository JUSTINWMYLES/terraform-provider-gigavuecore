package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadPermittedAddressesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadPermittedAddressesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadPermittedAddressesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadPermittedAddressesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadPermittedAddressesDataSourceMetadata(t *testing.T) {
	d := NewLoadPermittedAddressesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_permitted_addresses" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_permitted_addresses")
	}
}
