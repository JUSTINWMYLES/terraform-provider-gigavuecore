package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSpineLinkAllDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadSpineLinkAllDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadSpineLinkAllDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadSpineLinkAllDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadSpineLinkAllDataSourceMetadata(t *testing.T) {
	d := NewLoadSpineLinkAllDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_spine_link_all" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_spine_link_all")
	}
}
