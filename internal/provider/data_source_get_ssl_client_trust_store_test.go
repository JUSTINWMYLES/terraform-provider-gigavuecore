package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslClientTrustStoreDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSslClientTrustStoreDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSslClientTrustStoreDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSslClientTrustStoreDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSslClientTrustStoreDataSourceMetadata(t *testing.T) {
	d := NewGetSslClientTrustStoreDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ssl_client_trust_store" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ssl_client_trust_store")
	}
}
