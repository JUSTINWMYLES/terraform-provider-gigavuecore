package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetMapChainDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetMapChainDataSourceSchemaValidation(t *testing.T) {
	d := NewGetMapChainDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetMapChainDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetMapChainDataSourceMetadata(t *testing.T) {
	d := NewGetMapChainDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_map_chain" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_map_chain")
	}
}
