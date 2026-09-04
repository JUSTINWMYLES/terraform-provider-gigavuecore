package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllCircuitTunnelGlobalDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllCircuitTunnelGlobalDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllCircuitTunnelGlobalDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllCircuitTunnelGlobalDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllCircuitTunnelGlobalDataSourceMetadata(t *testing.T) {
	d := NewLoadAllCircuitTunnelGlobalDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_circuit_tunnel_global" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_circuit_tunnel_global")
	}
}
