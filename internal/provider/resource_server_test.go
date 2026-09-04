package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestServerResourceSchemaValidation verifies that the generated resource schema is valid.
func TestServerResourceSchemaValidation(t *testing.T) {
	r := &ServerResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestServerResourceMetadata verifies that the generated resource reports the expected type name.
func TestServerResourceMetadata(t *testing.T) {
	r := &ServerResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_server" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_server")
	}
}
