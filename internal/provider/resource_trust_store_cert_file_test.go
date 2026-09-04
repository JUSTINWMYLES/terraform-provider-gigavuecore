package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestTrustStoreCertFileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestTrustStoreCertFileResourceSchemaValidation(t *testing.T) {
	r := &TrustStoreCertFileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestTrustStoreCertFileResourceMetadata verifies that the generated resource reports the expected type name.
func TestTrustStoreCertFileResourceMetadata(t *testing.T) {
	r := &TrustStoreCertFileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_trust_store_cert_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_trust_store_cert_file")
	}
}
