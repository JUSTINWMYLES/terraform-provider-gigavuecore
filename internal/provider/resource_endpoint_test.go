package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestEndpointResourceSchemaValidation verifies that the generated resource schema is valid.
func TestEndpointResourceSchemaValidation(t *testing.T) {
	r := &EndpointResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestEndpointResourceMetadata verifies that the generated resource reports the expected type name.
func TestEndpointResourceMetadata(t *testing.T) {
	r := &EndpointResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_endpoint" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_endpoint")
	}
}
