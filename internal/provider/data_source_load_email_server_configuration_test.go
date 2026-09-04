package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadEmailServerConfigurationDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadEmailServerConfigurationDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadEmailServerConfigurationDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadEmailServerConfigurationDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadEmailServerConfigurationDataSourceMetadata(t *testing.T) {
	d := NewLoadEmailServerConfigurationDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_email_server_configuration" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_email_server_configuration")
	}
}
