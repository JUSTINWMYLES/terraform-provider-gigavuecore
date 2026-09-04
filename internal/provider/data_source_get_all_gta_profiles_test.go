package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllGtaProfilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllGtaProfilesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllGtaProfilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllGtaProfilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllGtaProfilesDataSourceMetadata(t *testing.T) {
	d := NewGetAllGtaProfilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_gta_profiles" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_gta_profiles")
	}
}
