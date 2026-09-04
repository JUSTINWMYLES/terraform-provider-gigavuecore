package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFmTemplatesHierarchialConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllFmTemplatesHierarchialConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllFmTemplatesHierarchialConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllFmTemplatesHierarchialConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllFmTemplatesHierarchialConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadAllFmTemplatesHierarchialConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_fm_templates_hierarchial_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_fm_templates_hierarchial_config")
	}
}
