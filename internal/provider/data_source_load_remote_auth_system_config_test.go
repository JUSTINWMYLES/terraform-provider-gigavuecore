package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadRemoteAuthSystemConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadRemoteAuthSystemConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadRemoteAuthSystemConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadRemoteAuthSystemConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadRemoteAuthSystemConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadRemoteAuthSystemConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_remote_auth_system_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_remote_auth_system_config")
	}
}
