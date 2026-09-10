package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineToolGroupsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllInlineToolGroupsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllInlineToolGroupsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllInlineToolGroupsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllInlineToolGroupsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllInlineToolGroupsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_inline_tool_groups" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_inline_tool_groups")
	}
}
