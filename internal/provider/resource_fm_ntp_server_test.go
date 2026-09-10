package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestFmNtpServerResourceSchemaValidation verifies that the generated resource schema is valid.
func TestFmNtpServerResourceSchemaValidation(t *testing.T) {
	r := &FmNtpServerResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestFmNtpServerResourceMetadata verifies that the generated resource reports the expected type name.
func TestFmNtpServerResourceMetadata(t *testing.T) {
	r := &FmNtpServerResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_fm_ntp_server" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_fm_ntp_server")
	}
}
