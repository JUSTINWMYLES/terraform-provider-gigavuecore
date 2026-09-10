package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllGsopsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllGsopsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllGsopsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllGsopsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllGsopsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllGsopsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_gsops" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_gsops")
	}
}
