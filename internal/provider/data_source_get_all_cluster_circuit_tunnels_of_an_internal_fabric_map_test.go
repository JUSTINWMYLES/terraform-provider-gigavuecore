package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSourceMetadata(t *testing.T) {
	d := NewGetAllClusterCircuitTunnelsOfAnInternalFabricMapDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_cluster_circuit_tunnels_of_an_internal_fabric_map")
	}
}
