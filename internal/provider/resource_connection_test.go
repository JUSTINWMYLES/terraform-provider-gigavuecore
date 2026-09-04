package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestConnectionResourceSchemaValidation verifies that the generated resource schema is valid.
func TestConnectionResourceSchemaValidation(t *testing.T) {
	r := &ConnectionResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestConnectionResourceMetadata verifies that the generated resource reports the expected type name.
func TestConnectionResourceMetadata(t *testing.T) {
	r := &ConnectionResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_connection" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_connection")
	}
}
