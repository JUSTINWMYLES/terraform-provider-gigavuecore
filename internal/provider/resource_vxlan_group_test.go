package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestVxlanGroupResourceSchemaValidation verifies that the generated resource schema is valid.
func TestVxlanGroupResourceSchemaValidation(t *testing.T) {
	r := &VxlanGroupResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestVxlanGroupResourceMetadata verifies that the generated resource reports the expected type name.
func TestVxlanGroupResourceMetadata(t *testing.T) {
	r := &VxlanGroupResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_vxlan_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_vxlan_group")
	}
}
