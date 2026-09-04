package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllGsGroupsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllGsGroupsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllGsGroupsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllGsGroupsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllGsGroupsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllGsGroupsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_gs_groups" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_gs_groups")
	}
}
