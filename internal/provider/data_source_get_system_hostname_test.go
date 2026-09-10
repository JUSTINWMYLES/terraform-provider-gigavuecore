package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemHostnameDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSystemHostnameDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSystemHostnameDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSystemHostnameDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSystemHostnameDataSourceMetadata(t *testing.T) {
	d := NewGetSystemHostnameDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_system_hostname" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_system_hostname")
	}
}
