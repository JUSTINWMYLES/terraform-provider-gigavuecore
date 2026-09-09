package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTrafficFlowsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllTrafficFlowsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllTrafficFlowsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllTrafficFlowsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllTrafficFlowsDataSourceMetadata(t *testing.T) {
	d := NewGetAllTrafficFlowsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_traffic_flows" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_traffic_flows")
	}
}
