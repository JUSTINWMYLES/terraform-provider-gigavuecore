package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestToolPortMirrorResourceSchemaValidation verifies that the generated resource schema is valid.
func TestToolPortMirrorResourceSchemaValidation(t *testing.T) {
	r := &ToolPortMirrorResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestToolPortMirrorResourceMetadata verifies that the generated resource reports the expected type name.
func TestToolPortMirrorResourceMetadata(t *testing.T) {
	r := &ToolPortMirrorResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_tool_port_mirror" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_tool_port_mirror")
	}
}
