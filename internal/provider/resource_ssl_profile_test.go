package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSslProfileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestSslProfileResourceSchemaValidation(t *testing.T) {
	r := &SslProfileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSslProfileResourceMetadata verifies that the generated resource reports the expected type name.
func TestSslProfileResourceMetadata(t *testing.T) {
	r := &SslProfileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_ssl_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_ssl_profile")
	}
}
