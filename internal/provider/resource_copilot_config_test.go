package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestCopilotConfigResourceSchemaValidation verifies that the generated resource schema is valid.
func TestCopilotConfigResourceSchemaValidation(t *testing.T) {
	r := &CopilotConfigResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestCopilotConfigResourceMetadata verifies that the generated resource reports the expected type name.
func TestCopilotConfigResourceMetadata(t *testing.T) {
	r := &CopilotConfigResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_copilot_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_copilot_config")
	}
}
