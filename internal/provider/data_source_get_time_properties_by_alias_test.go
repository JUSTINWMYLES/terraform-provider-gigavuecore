package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTimePropertiesByAliasDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetTimePropertiesByAliasDataSourceSchemaValidation(t *testing.T) {
	d := NewGetTimePropertiesByAliasDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetTimePropertiesByAliasDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetTimePropertiesByAliasDataSourceMetadata(t *testing.T) {
	d := NewGetTimePropertiesByAliasDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_time_properties_by_alias" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_time_properties_by_alias")
	}
}
