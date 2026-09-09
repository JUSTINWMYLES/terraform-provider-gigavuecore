package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFlexInlineSslAppsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllFlexInlineSslAppsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllFlexInlineSslAppsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllFlexInlineSslAppsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllFlexInlineSslAppsDataSourceMetadata(t *testing.T) {
	d := NewGetAllFlexInlineSslAppsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_flex_inline_ssl_apps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_flex_inline_ssl_apps")
	}
}
