package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllMapTemplatesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllMapTemplatesDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllMapTemplatesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllMapTemplatesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllMapTemplatesDataSourceMetadata(t *testing.T) {
	d := NewLoadAllMapTemplatesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_map_templates" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_map_templates")
	}
}
