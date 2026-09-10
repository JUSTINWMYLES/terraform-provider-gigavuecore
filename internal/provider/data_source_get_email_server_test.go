package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEmailServerDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetEmailServerDataSourceSchemaValidation(t *testing.T) {
	d := NewGetEmailServerDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetEmailServerDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetEmailServerDataSourceMetadata(t *testing.T) {
	d := NewGetEmailServerDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_email_server" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_email_server")
	}
}
