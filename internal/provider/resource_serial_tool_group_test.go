package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestSerialToolGroupResourceSchemaValidation verifies that the generated resource schema is valid.
func TestSerialToolGroupResourceSchemaValidation(t *testing.T) {
	r := &SerialToolGroupResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSerialToolGroupResourceMetadata verifies that the generated resource reports the expected type name.
func TestSerialToolGroupResourceMetadata(t *testing.T) {
	r := &SerialToolGroupResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_serial_tool_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_serial_tool_group")
	}
}
