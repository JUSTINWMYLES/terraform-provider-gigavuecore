package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestConfigFileResourceSchemaValidation verifies that the generated resource schema is valid.
func TestConfigFileResourceSchemaValidation(t *testing.T) {
	r := &ConfigFileResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestConfigFileResourceMetadata verifies that the generated resource reports the expected type name.
func TestConfigFileResourceMetadata(t *testing.T) {
	r := &ConfigFileResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_config_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_config_file")
	}
}
