package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadInlineSslUrlStatusDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadInlineSslUrlStatusDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadInlineSslUrlStatusDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadInlineSslUrlStatusDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadInlineSslUrlStatusDataSourceMetadata(t *testing.T) {
	d := NewLoadInlineSslUrlStatusDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_inline_ssl_url_status" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_inline_ssl_url_status")
	}
}
