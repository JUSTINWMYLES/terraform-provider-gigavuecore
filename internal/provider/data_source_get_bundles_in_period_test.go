package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetBundlesInPeriodDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetBundlesInPeriodDataSourceSchemaValidation(t *testing.T) {
	d := NewGetBundlesInPeriodDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetBundlesInPeriodDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetBundlesInPeriodDataSourceMetadata(t *testing.T) {
	d := NewGetBundlesInPeriodDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_bundles_in_period" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_bundles_in_period")
	}
}
