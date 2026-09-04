package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestAdvHashResourceSchemaValidation verifies that the generated resource schema is valid.
func TestAdvHashResourceSchemaValidation(t *testing.T) {
	r := &AdvHashResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestAdvHashResourceMetadata verifies that the generated resource reports the expected type name.
func TestAdvHashResourceMetadata(t *testing.T) {
	r := &AdvHashResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_adv_hash" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_adv_hash")
	}
}
