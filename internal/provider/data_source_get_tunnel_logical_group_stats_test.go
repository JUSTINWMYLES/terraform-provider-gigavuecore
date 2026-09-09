package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTunnelLogicalGroupStatsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetTunnelLogicalGroupStatsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetTunnelLogicalGroupStatsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetTunnelLogicalGroupStatsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetTunnelLogicalGroupStatsDataSourceMetadata(t *testing.T) {
	d := NewGetTunnelLogicalGroupStatsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_tunnel_logical_group_stats" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_tunnel_logical_group_stats")
	}
}
