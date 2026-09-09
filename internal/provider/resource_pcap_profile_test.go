package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestPcapProfileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestPcapProfileResourceSchemaValidation(t *testing.T) {
	r := &PcapProfileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestPcapProfileResourceMetadata verifies that the generated resource reports the expected type name.
func TestPcapProfileResourceMetadata(t *testing.T) {
	r := &PcapProfileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_pcap_profile" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_pcap_profile")
	}
}
