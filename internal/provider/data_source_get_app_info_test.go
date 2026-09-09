package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAppInfoDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAppInfoDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAppInfoDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAppInfoDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAppInfoDataSourceMetadata(t *testing.T) {
	d := NewGetAppInfoDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_app_info" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_app_info")
	}
}
