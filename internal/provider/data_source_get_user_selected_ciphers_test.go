package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUserSelectedCiphersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetUserSelectedCiphersDataSourceSchemaValidation(t *testing.T) {
	d := NewGetUserSelectedCiphersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetUserSelectedCiphersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetUserSelectedCiphersDataSourceMetadata(t *testing.T) {
	d := NewGetUserSelectedCiphersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_user_selected_ciphers" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_user_selected_ciphers")
	}
}
