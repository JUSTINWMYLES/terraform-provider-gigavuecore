package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEmsPreferencesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetEmsPreferencesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetEmsPreferencesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetEmsPreferencesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetEmsPreferencesDataSourceMetadata(t *testing.T) {
	d := NewGetEmsPreferencesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ems_preferences" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ems_preferences")
	}
}
