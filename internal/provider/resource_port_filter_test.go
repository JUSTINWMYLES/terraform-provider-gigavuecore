package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPortFilterResourceSchemaValidation verifies that the generated resource schema is valid.
func TestPortFilterResourceSchemaValidation(t *testing.T) {
	r := &PortFilterResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestPortFilterResourceMetadata verifies that the generated resource reports the expected type name.
func TestPortFilterResourceMetadata(t *testing.T) {
	r := &PortFilterResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_port_filter" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_port_filter")
	}
}
