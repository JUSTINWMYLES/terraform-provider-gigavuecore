package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPolicyNrtStatsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPolicyNrtStatsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPolicyNrtStatsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPolicyNrtStatsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPolicyNrtStatsDataSourceMetadata(t *testing.T) {
	d := NewGetPolicyNrtStatsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_policy_nrt_stats" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_policy_nrt_stats")
	}
}
