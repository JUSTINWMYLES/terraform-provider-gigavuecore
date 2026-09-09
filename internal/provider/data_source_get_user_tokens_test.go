package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetUserTokensDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetUserTokensDataSourceSchemaValidation(t *testing.T) {
	d := NewGetUserTokensDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetUserTokensDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetUserTokensDataSourceMetadata(t *testing.T) {
	d := NewGetUserTokensDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_user_tokens" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_user_tokens")
	}
}
