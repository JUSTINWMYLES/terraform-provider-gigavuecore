package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNodeResourceSchemaValidation verifies that the generated resource schema is valid.
func TestNodeResourceSchemaValidation(t *testing.T) {
	r := &NodeResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestNodeResourceMetadata verifies that the generated resource reports the expected type name.
func TestNodeResourceMetadata(t *testing.T) {
	r := &NodeResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_node" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_node")
	}
}
