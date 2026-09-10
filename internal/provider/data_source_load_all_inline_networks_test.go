package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineNetworksDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllInlineNetworksDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllInlineNetworksDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllInlineNetworksDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllInlineNetworksDataSourceMetadata(t *testing.T) {
	d := NewLoadAllInlineNetworksDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_inline_networks" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_inline_networks")
	}
}
