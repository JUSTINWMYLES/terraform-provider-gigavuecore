package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetInlineSslConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetInlineSslConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewGetInlineSslConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetInlineSslConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetInlineSslConfigDataSourceMetadata(t *testing.T) {
	d := NewGetInlineSslConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_inline_ssl_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_inline_ssl_config")
	}
}
