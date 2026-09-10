package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSaApfProfileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestSaApfProfileResourceSchemaValidation(t *testing.T) {
	r := &SaApfProfileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSaApfProfileResourceMetadata verifies that the generated resource reports the expected type name.
func TestSaApfProfileResourceMetadata(t *testing.T) {
	r := &SaApfProfileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_sa_apf_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_sa_apf_profile")
	}
}
