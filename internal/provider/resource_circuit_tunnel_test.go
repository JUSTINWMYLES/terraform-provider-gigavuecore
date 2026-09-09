package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestCircuitTunnelResourceSchemaValidation verifies that the generated resource schema is valid.
func TestCircuitTunnelResourceSchemaValidation(t *testing.T) {
	r := &CircuitTunnelResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestCircuitTunnelResourceMetadata verifies that the generated resource reports the expected type name.
func TestCircuitTunnelResourceMetadata(t *testing.T) {
	r := &CircuitTunnelResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_circuit_tunnel" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_circuit_tunnel")
	}
}
