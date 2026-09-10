package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTunnelLogicalGroupsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllTunnelLogicalGroupsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllTunnelLogicalGroupsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllTunnelLogicalGroupsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllTunnelLogicalGroupsDataSourceMetadata(t *testing.T) {
	d := NewGetAllTunnelLogicalGroupsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_tunnel_logical_groups" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_tunnel_logical_groups")
	}
}
