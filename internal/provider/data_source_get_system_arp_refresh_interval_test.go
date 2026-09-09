package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemArpRefreshIntervalDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSystemArpRefreshIntervalDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSystemArpRefreshIntervalDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSystemArpRefreshIntervalDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSystemArpRefreshIntervalDataSourceMetadata(t *testing.T) {
	d := NewGetSystemArpRefreshIntervalDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_system_arp_refresh_interval" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_system_arp_refresh_interval")
	}
}
