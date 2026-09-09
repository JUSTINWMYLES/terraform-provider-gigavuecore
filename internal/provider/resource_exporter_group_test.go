package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestExporterGroupResourceSchemaValidation verifies that the generated resource schema is valid.
func TestExporterGroupResourceSchemaValidation(t *testing.T) {
	r := &ExporterGroupResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestExporterGroupResourceMetadata verifies that the generated resource reports the expected type name.
func TestExporterGroupResourceMetadata(t *testing.T) {
	r := &ExporterGroupResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_exporter_group" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_exporter_group")
	}
}
