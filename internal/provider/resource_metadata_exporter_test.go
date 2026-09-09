package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMetadataExporterResourceSchemaValidation verifies that the generated resource schema is valid.
func TestMetadataExporterResourceSchemaValidation(t *testing.T) {
	r := &MetadataExporterResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestMetadataExporterResourceMetadata verifies that the generated resource reports the expected type name.
func TestMetadataExporterResourceMetadata(t *testing.T) {
	r := &MetadataExporterResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_metadata_exporter" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_metadata_exporter")
	}
}
