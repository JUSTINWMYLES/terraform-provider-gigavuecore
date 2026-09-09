package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllFilterTemplatesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllFilterTemplatesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllFilterTemplatesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllFilterTemplatesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllFilterTemplatesDataSourceMetadata(t *testing.T) {
	d := NewLoadAllFilterTemplatesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_filter_templates" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_filter_templates")
	}
}
