package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllUserTagsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllUserTagsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllUserTagsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllUserTagsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllUserTagsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllUserTagsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_user_tags" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_user_tags")
	}
}
