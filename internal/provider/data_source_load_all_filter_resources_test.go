package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFilterResourcesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllFilterResourcesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllFilterResourcesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllFilterResourcesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllFilterResourcesDataSourceMetadata(t *testing.T) {
	d := NewLoadAllFilterResourcesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_filter_resources" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_filter_resources")
	}
}
