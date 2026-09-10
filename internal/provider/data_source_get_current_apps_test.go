package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentAppsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCurrentAppsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCurrentAppsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCurrentAppsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCurrentAppsDataSourceMetadata(t *testing.T) {
	d := NewGetCurrentAppsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_current_apps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_current_apps")
	}
}
