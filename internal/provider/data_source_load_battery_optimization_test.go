package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadBatteryOptimizationDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadBatteryOptimizationDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadBatteryOptimizationDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadBatteryOptimizationDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadBatteryOptimizationDataSourceMetadata(t *testing.T) {
	d := NewLoadBatteryOptimizationDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_battery_optimization" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_battery_optimization")
	}
}
