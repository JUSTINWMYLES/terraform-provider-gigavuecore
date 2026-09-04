package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadFmLicensingSummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadFmLicensingSummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadFmLicensingSummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadFmLicensingSummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadFmLicensingSummaryDataSourceMetadata(t *testing.T) {
	d := NewLoadFmLicensingSummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_fm_licensing_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_fm_licensing_summary")
	}
}
