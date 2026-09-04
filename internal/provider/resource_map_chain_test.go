package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMapChainResourceSchemaValidation verifies that the generated resource schema is valid.
func TestMapChainResourceSchemaValidation(t *testing.T) {
	r := &MapChainResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestMapChainResourceMetadata verifies that the generated resource reports the expected type name.
func TestMapChainResourceMetadata(t *testing.T) {
	r := &MapChainResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_map_chain" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_map_chain")
	}
}
