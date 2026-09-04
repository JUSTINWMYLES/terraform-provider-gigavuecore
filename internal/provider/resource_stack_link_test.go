package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestStackLinkResourceSchemaValidation verifies that the generated resource schema is valid.
func TestStackLinkResourceSchemaValidation(t *testing.T) {
	r := &StackLinkResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestStackLinkResourceMetadata verifies that the generated resource reports the expected type name.
func TestStackLinkResourceMetadata(t *testing.T) {
	r := &StackLinkResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_stack_link" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_stack_link")
	}
}
