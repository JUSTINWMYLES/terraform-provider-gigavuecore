package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFmLicensesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadFmLicensesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadFmLicensesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadFmLicensesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadFmLicensesDataSourceMetadata(t *testing.T) {
	d := NewLoadFmLicensesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_fm_licenses" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_fm_licenses")
	}
}
