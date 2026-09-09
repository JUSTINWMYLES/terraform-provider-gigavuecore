package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNodeCredentialsResourceSchemaValidation verifies that the generated resource schema is valid.
func TestNodeCredentialsResourceSchemaValidation(t *testing.T) {
	r := &NodeCredentialsResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestNodeCredentialsResourceMetadata verifies that the generated resource reports the expected type name.
func TestNodeCredentialsResourceMetadata(t *testing.T) {
	r := &NodeCredentialsResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_node_credentials" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_node_credentials")
	}
}
