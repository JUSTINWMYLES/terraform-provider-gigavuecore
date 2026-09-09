package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCurrentMonSessionsToAppTierMapDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCurrentMonSessionsToAppTierMapDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCurrentMonSessionsToAppTierMapDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCurrentMonSessionsToAppTierMapDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCurrentMonSessionsToAppTierMapDataSourceMetadata(t *testing.T) {
	d := NewGetCurrentMonSessionsToAppTierMapDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_current_mon_sessions_to_app_tier_map" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_current_mon_sessions_to_app_tier_map")
	}
}
