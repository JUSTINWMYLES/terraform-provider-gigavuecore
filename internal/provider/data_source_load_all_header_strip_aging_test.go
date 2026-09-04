package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllHeaderStripAgingDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllHeaderStripAgingDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllHeaderStripAgingDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllHeaderStripAgingDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllHeaderStripAgingDataSourceMetadata(t *testing.T) {
	d := NewLoadAllHeaderStripAgingDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_header_strip_aging" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_header_strip_aging")
	}
}
