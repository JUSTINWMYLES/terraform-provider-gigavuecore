package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTcpProfilesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllTcpProfilesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllTcpProfilesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllTcpProfilesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllTcpProfilesDataSourceMetadata(t *testing.T) {
	d := NewGetAllTcpProfilesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_tcp_profiles" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_tcp_profiles")
	}
}
