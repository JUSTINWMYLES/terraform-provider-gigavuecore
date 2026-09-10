package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTrafficPolicyGraphDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllTrafficPolicyGraphDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllTrafficPolicyGraphDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllTrafficPolicyGraphDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllTrafficPolicyGraphDataSourceMetadata(t *testing.T) {
	d := NewGetAllTrafficPolicyGraphDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_traffic_policy_graph" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_traffic_policy_graph")
	}
}
