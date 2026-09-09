package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSubFlowResourceSchemaValidation verifies that the generated resource schema is valid.
func TestSubFlowResourceSchemaValidation(t *testing.T) {
	r := &SubFlowResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSubFlowResourceMetadata verifies that the generated resource reports the expected type name.
func TestSubFlowResourceMetadata(t *testing.T) {
	r := &SubFlowResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_sub_flow" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_sub_flow")
	}
}
