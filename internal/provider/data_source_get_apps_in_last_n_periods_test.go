package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAppsInLastNPeriodsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAppsInLastNPeriodsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAppsInLastNPeriodsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAppsInLastNPeriodsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAppsInLastNPeriodsDataSourceMetadata(t *testing.T) {
	d := NewGetAppsInLastNPeriodsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_apps_in_last_n_periods" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_apps_in_last_n_periods")
	}
}
