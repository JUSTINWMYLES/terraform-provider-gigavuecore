package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllEnhancedSlicingsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllEnhancedSlicingsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllEnhancedSlicingsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllEnhancedSlicingsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllEnhancedSlicingsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllEnhancedSlicingsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_enhanced_slicings" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_enhanced_slicings")
	}
}
