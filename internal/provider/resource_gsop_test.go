package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGsopResourceSchemaValidation verifies that the generated resource schema is valid.
func TestGsopResourceSchemaValidation(t *testing.T) {
	r := &GsopResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGsopResourceMetadata verifies that the generated resource reports the expected type name.
func TestGsopResourceMetadata(t *testing.T) {
	r := &GsopResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_gsop" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_gsop")
	}
}
