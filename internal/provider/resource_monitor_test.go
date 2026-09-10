package provider

import (
	"context"
	"testing"
)
import tfframeworkresource "github.com/hashicorp/terraform-plugin-framework/resource"

// TestMonitorResourceSchemaValidation verifies that the generated resource schema is valid.
func TestMonitorResourceSchemaValidation(t *testing.T) {
	r := &MonitorResource{}
	var resp tfframeworkresource.SchemaResponse
	r.Schema(context.Background(), tfframeworkresource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestMonitorResourceMetadata verifies that the generated resource reports the expected type name.
func TestMonitorResourceMetadata(t *testing.T) {
	r := &MonitorResource{}
	var resp tfframeworkresource.MetadataResponse
	r.Metadata(context.Background(), tfframeworkresource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_monitor" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_monitor")
	}
}
