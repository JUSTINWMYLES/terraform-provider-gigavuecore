package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestGtpWhitelistResourceSchemaValidation verifies that the generated resource schema is valid.
func TestGtpWhitelistResourceSchemaValidation(t *testing.T) {
	r := &GtpWhitelistResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGtpWhitelistResourceMetadata verifies that the generated resource reports the expected type name.
func TestGtpWhitelistResourceMetadata(t *testing.T) {
	r := &GtpWhitelistResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_gtp_whitelist" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_gtp_whitelist")
	}
}
