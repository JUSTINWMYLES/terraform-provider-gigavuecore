package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllSslProfilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllSslProfilesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllSslProfilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllSslProfilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllSslProfilesDataSourceMetadata(t *testing.T) {
	d := NewGetAllSslProfilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_ssl_profiles" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_ssl_profiles")
	}
}
