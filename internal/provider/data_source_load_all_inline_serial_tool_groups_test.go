package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllInlineSerialToolGroupsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllInlineSerialToolGroupsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllInlineSerialToolGroupsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllInlineSerialToolGroupsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllInlineSerialToolGroupsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllInlineSerialToolGroupsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_inline_serial_tool_groups" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_inline_serial_tool_groups")
	}
}
