package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetDeviceLocatorLedDetailsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetDeviceLocatorLedDetailsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetDeviceLocatorLedDetailsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetDeviceLocatorLedDetailsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetDeviceLocatorLedDetailsDataSourceMetadata(t *testing.T) {
	d := NewGetDeviceLocatorLedDetailsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_device_locator_led_details" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_device_locator_led_details")
	}
}
