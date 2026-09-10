package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTopovizExternalDevicesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetTopovizExternalDevicesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetTopovizExternalDevicesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetTopovizExternalDevicesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetTopovizExternalDevicesDataSourceMetadata(t *testing.T) {
	d := NewGetTopovizExternalDevicesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_topoviz_external_devices" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_topoviz_external_devices")
	}
}
