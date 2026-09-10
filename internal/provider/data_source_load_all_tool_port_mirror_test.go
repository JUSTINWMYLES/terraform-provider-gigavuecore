package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllToolPortMirrorDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllToolPortMirrorDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllToolPortMirrorDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllToolPortMirrorDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllToolPortMirrorDataSourceMetadata(t *testing.T) {
	d := NewLoadAllToolPortMirrorDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_tool_port_mirror" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_tool_port_mirror")
	}
}
