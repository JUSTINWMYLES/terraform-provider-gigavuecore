package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSipWhitelistResourceSchemaValidation verifies that the generated resource schema is valid.
func TestSipWhitelistResourceSchemaValidation(t *testing.T) {
	r := &SipWhitelistResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSipWhitelistResourceMetadata verifies that the generated resource reports the expected type name.
func TestSipWhitelistResourceMetadata(t *testing.T) {
	r := &SipWhitelistResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_sip_whitelist" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_sip_whitelist")
	}
}
