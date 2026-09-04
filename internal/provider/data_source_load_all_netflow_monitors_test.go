package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllNetflowMonitorsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllNetflowMonitorsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllNetflowMonitorsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllNetflowMonitorsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllNetflowMonitorsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllNetflowMonitorsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_netflow_monitors" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_netflow_monitors")
	}
}
