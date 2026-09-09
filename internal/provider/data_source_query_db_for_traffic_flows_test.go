package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestQueryDbForTrafficFlowsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestQueryDbForTrafficFlowsDataSourceSchemaValidation(t *testing.T) {
	d := NewQueryDbForTrafficFlowsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestQueryDbForTrafficFlowsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestQueryDbForTrafficFlowsDataSourceMetadata(t *testing.T) {
	d := NewQueryDbForTrafficFlowsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_query_db_for_traffic_flows" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_query_db_for_traffic_flows")
	}
}
