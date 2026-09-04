package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadInlineSslCertValidStatusDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadInlineSslCertValidStatusDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadInlineSslCertValidStatusDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadInlineSslCertValidStatusDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadInlineSslCertValidStatusDataSourceMetadata(t *testing.T) {
	d := NewLoadInlineSslCertValidStatusDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_inline_ssl_cert_valid_status" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_inline_ssl_cert_valid_status")
	}
}
