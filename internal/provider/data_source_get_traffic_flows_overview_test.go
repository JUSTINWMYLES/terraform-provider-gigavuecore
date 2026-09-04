package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrafficFlowsOverviewDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetTrafficFlowsOverviewDataSourceSchemaValidation(t *testing.T) {
	d := NewGetTrafficFlowsOverviewDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetTrafficFlowsOverviewDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetTrafficFlowsOverviewDataSourceMetadata(t *testing.T) {
	d := NewGetTrafficFlowsOverviewDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_traffic_flows_overview" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_traffic_flows_overview")
	}
}
