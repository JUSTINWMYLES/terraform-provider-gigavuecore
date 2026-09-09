package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrafficFlowsGlobalSettingsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetTrafficFlowsGlobalSettingsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetTrafficFlowsGlobalSettingsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetTrafficFlowsGlobalSettingsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetTrafficFlowsGlobalSettingsDataSourceMetadata(t *testing.T) {
	d := NewGetTrafficFlowsGlobalSettingsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_traffic_flows_global_settings" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_traffic_flows_global_settings")
	}
}
