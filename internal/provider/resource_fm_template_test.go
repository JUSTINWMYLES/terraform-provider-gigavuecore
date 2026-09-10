package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFmTemplateResourceSchemaValidation verifies that the generated resource schema is valid.
func TestFmTemplateResourceSchemaValidation(t *testing.T) {
	r := &FmTemplateResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestFmTemplateResourceMetadata verifies that the generated resource reports the expected type name.
func TestFmTemplateResourceMetadata(t *testing.T) {
	r := &FmTemplateResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_fm_template" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_fm_template")
	}
}
