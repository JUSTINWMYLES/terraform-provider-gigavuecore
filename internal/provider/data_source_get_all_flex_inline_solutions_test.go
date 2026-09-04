package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllFlexInlineSolutionsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllFlexInlineSolutionsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllFlexInlineSolutionsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllFlexInlineSolutionsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllFlexInlineSolutionsDataSourceMetadata(t *testing.T) {
	d := NewGetAllFlexInlineSolutionsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_flex_inline_solutions" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_flex_inline_solutions")
	}
}
