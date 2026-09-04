package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslClientTrustStoreFileDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSslClientTrustStoreFileDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSslClientTrustStoreFileDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSslClientTrustStoreFileDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSslClientTrustStoreFileDataSourceMetadata(t *testing.T) {
	d := NewGetSslClientTrustStoreFileDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ssl_client_trust_store_file" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ssl_client_trust_store_file")
	}
}
