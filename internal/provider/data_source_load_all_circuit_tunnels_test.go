package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllCircuitTunnelsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllCircuitTunnelsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllCircuitTunnelsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllCircuitTunnelsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllCircuitTunnelsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllCircuitTunnelsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_circuit_tunnels" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_circuit_tunnels")
	}
}
