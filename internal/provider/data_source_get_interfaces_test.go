package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetInterfacesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetInterfacesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetInterfacesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetInterfacesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetInterfacesDataSourceMetadata(t *testing.T) {
	d := NewGetInterfacesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_interfaces" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_interfaces")
	}
}
