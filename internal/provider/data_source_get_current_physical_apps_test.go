package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentPhysicalAppsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCurrentPhysicalAppsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCurrentPhysicalAppsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCurrentPhysicalAppsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCurrentPhysicalAppsDataSourceMetadata(t *testing.T) {
	d := NewGetCurrentPhysicalAppsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_current_physical_apps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_current_physical_apps")
	}
}
