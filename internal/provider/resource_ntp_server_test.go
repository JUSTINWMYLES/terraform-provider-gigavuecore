package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNtpServerResourceSchemaValidation verifies that the generated resource schema is valid.
func TestNtpServerResourceSchemaValidation(t *testing.T) {
	r := &NtpServerResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestNtpServerResourceMetadata verifies that the generated resource reports the expected type name.
func TestNtpServerResourceMetadata(t *testing.T) {
	r := &NtpServerResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_ntp_server" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_ntp_server")
	}
}
