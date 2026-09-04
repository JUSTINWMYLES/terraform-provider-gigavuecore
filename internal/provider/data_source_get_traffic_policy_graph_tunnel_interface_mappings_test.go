package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetTrafficPolicyGraphTunnelInterfaceMappingsDataSourceMetadata(t *testing.T) {
	d := NewGetTrafficPolicyGraphTunnelInterfaceMappingsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_traffic_policy_graph_tunnel_interface_mappings")
	}
}
