package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestIcapResourceSchemaValidation verifies that the generated resource schema is valid.
func TestIcapResourceSchemaValidation(t *testing.T) {
	r := &IcapResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestIcapResourceMetadata verifies that the generated resource reports the expected type name.
func TestIcapResourceMetadata(t *testing.T) {
	r := &IcapResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_icap" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_icap")
	}
}
