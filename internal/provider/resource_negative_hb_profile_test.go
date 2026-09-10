package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNegativeHbProfileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestNegativeHbProfileResourceSchemaValidation(t *testing.T) {
	r := &NegativeHbProfileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestNegativeHbProfileResourceMetadata verifies that the generated resource reports the expected type name.
func TestNegativeHbProfileResourceMetadata(t *testing.T) {
	r := &NegativeHbProfileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_negative_hb_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_negative_hb_profile")
	}
}
