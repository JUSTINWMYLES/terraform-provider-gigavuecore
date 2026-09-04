package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestRadiusServerResourceSchemaValidation verifies that the generated resource schema is valid.
func TestRadiusServerResourceSchemaValidation(t *testing.T) {
	r := &RadiusServerResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestRadiusServerResourceMetadata verifies that the generated resource reports the expected type name.
func TestRadiusServerResourceMetadata(t *testing.T) {
	r := &RadiusServerResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_radius_server" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_radius_server")
	}
}
