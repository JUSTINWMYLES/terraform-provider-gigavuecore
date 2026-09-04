package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslClientTrustStoreListDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSslClientTrustStoreListDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSslClientTrustStoreListDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSslClientTrustStoreListDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSslClientTrustStoreListDataSourceMetadata(t *testing.T) {
	d := NewGetSslClientTrustStoreListDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ssl_client_trust_store_list" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ssl_client_trust_store_list")
	}
}
