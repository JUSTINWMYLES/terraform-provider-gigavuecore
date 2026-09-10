package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestElbResourceSchemaValidation verifies that the generated resource schema is valid.
func TestElbResourceSchemaValidation(t *testing.T) {
	r := &ElbResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestElbResourceMetadata verifies that the generated resource reports the expected type name.
func TestElbResourceMetadata(t *testing.T) {
	r := &ElbResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_elb" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_elb")
	}
}
