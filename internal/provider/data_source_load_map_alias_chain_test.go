package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadMapAliasChainDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadMapAliasChainDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadMapAliasChainDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadMapAliasChainDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadMapAliasChainDataSourceMetadata(t *testing.T) {
	d := NewLoadMapAliasChainDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_map_alias_chain" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_map_alias_chain")
	}
}
