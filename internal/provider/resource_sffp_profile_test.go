package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSffpProfileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestSffpProfileResourceSchemaValidation(t *testing.T) {
	r := &SffpProfileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSffpProfileResourceMetadata verifies that the generated resource reports the expected type name.
func TestSffpProfileResourceMetadata(t *testing.T) {
	r := &SffpProfileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_sffp_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_sffp_profile")
	}
}
