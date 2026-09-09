package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestInlineNetworkGroupResourceSchemaValidation verifies that the generated resource schema is valid.
func TestInlineNetworkGroupResourceSchemaValidation(t *testing.T) {
	r := &InlineNetworkGroupResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestInlineNetworkGroupResourceMetadata verifies that the generated resource reports the expected type name.
func TestInlineNetworkGroupResourceMetadata(t *testing.T) {
	r := &InlineNetworkGroupResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_inline_network_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_inline_network_group")
	}
}
