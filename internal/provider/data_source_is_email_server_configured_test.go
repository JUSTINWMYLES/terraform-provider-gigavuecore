package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestIsEmailServerConfiguredDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestIsEmailServerConfiguredDataSourceSchemaValidation(t *testing.T) {
	d := NewIsEmailServerConfiguredDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestIsEmailServerConfiguredDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestIsEmailServerConfiguredDataSourceMetadata(t *testing.T) {
	d := NewIsEmailServerConfiguredDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_is_email_server_configured" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_is_email_server_configured")
	}
}
