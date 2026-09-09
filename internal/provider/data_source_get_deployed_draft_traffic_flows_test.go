package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetDeployedDraftTrafficFlowsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetDeployedDraftTrafficFlowsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetDeployedDraftTrafficFlowsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetDeployedDraftTrafficFlowsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetDeployedDraftTrafficFlowsDataSourceMetadata(t *testing.T) {
	d := NewGetDeployedDraftTrafficFlowsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_deployed_draft_traffic_flows" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_deployed_draft_traffic_flows")
	}
}
