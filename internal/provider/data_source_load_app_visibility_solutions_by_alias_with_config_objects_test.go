package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSourceMetadata(t *testing.T) {
	d := NewLoadAppVisibilitySolutionsByAliasWithConfigObjectsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_app_visibility_solutions_by_alias_with_config_objects" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_app_visibility_solutions_by_alias_with_config_objects")
	}
}
