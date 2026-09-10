package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemHostBannerDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSystemHostBannerDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSystemHostBannerDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSystemHostBannerDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSystemHostBannerDataSourceMetadata(t *testing.T) {
	d := NewGetSystemHostBannerDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_system_host_banner" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_system_host_banner")
	}
}
