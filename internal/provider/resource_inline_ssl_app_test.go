package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestInlineSslAppResourceSchemaValidation verifies that the generated resource schema is valid.
func TestInlineSslAppResourceSchemaValidation(t *testing.T) {
	r := &InlineSslAppResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestInlineSslAppResourceMetadata verifies that the generated resource reports the expected type name.
func TestInlineSslAppResourceMetadata(t *testing.T) {
	r := &InlineSslAppResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_inline_ssl_app" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_inline_ssl_app")
	}
}
