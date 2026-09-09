package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetOverlapComponentsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetOverlapComponentsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetOverlapComponentsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetOverlapComponentsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetOverlapComponentsDataSourceMetadata(t *testing.T) {
	d := NewGetOverlapComponentsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_overlap_components" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_overlap_components")
	}
}
