package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFilterTemplateResourceSchemaValidation verifies that the generated resource schema is valid.
func TestFilterTemplateResourceSchemaValidation(t *testing.T) {
	r := &FilterTemplateResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestFilterTemplateResourceMetadata verifies that the generated resource reports the expected type name.
func TestFilterTemplateResourceMetadata(t *testing.T) {
	r := &FilterTemplateResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_filter_template" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_filter_template")
	}
}
