package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllPlatformFilterLimitsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllPlatformFilterLimitsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllPlatformFilterLimitsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllPlatformFilterLimitsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllPlatformFilterLimitsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllPlatformFilterLimitsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_platform_filter_limits" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_platform_filter_limits")
	}
}
