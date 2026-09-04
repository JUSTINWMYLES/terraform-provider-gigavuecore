package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestServerGroupResourceSchemaValidation verifies that the generated resource schema is valid.
func TestServerGroupResourceSchemaValidation(t *testing.T) {
	r := &ServerGroupResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestServerGroupResourceMetadata verifies that the generated resource reports the expected type name.
func TestServerGroupResourceMetadata(t *testing.T) {
	r := &ServerGroupResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_server_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_server_group")
	}
}
