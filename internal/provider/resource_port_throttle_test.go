package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPortThrottleResourceSchemaValidation verifies that the generated resource schema is valid.
func TestPortThrottleResourceSchemaValidation(t *testing.T) {
	r := &PortThrottleResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestPortThrottleResourceMetadata verifies that the generated resource reports the expected type name.
func TestPortThrottleResourceMetadata(t *testing.T) {
	r := &PortThrottleResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_port_throttle" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_port_throttle")
	}
}
