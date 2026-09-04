package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllProxyServerProfilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllProxyServerProfilesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllProxyServerProfilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllProxyServerProfilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllProxyServerProfilesDataSourceMetadata(t *testing.T) {
	d := NewGetAllProxyServerProfilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_proxy_server_profiles" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_proxy_server_profiles")
	}
}
