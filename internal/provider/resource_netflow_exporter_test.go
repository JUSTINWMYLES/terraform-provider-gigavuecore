package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestNetflowExporterResourceSchemaValidation verifies that the generated resource schema is valid.
func TestNetflowExporterResourceSchemaValidation(t *testing.T) {
	r := &NetflowExporterResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestNetflowExporterResourceMetadata verifies that the generated resource reports the expected type name.
func TestNetflowExporterResourceMetadata(t *testing.T) {
	r := &NetflowExporterResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_netflow_exporter" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_netflow_exporter")
	}
}
