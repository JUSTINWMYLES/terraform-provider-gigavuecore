package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSourceSchemaValidation(t *testing.T) {
	d := NewGetClusterCircuitTunnelOfAnInternalFabricMapDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetClusterCircuitTunnelOfAnInternalFabricMapDataSourceMetadata(t *testing.T) {
	d := NewGetClusterCircuitTunnelOfAnInternalFabricMapDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_cluster_circuit_tunnel_of_an_internal_fabric_map" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_cluster_circuit_tunnel_of_an_internal_fabric_map")
	}
}
