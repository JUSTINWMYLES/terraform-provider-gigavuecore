package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslTrustStoreFileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSslTrustStoreFileDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSslTrustStoreFileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSslTrustStoreFileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSslTrustStoreFileDataSourceMetadata(t *testing.T) {
	d := NewGetSslTrustStoreFileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ssl_trust_store_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ssl_trust_store_file")
	}
}
