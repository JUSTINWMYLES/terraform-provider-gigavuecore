package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetLicenseSkusDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetLicenseSkusDataSourceSchemaValidation(t *testing.T) {
	d := NewGetLicenseSkusDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetLicenseSkusDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetLicenseSkusDataSourceMetadata(t *testing.T) {
	d := NewGetLicenseSkusDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_license_skus" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_license_skus")
	}
}
