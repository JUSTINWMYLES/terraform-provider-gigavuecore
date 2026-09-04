package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrustStoreListDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetTrustStoreListDataSourceSchemaValidation(t *testing.T) {
	d := NewGetTrustStoreListDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetTrustStoreListDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetTrustStoreListDataSourceMetadata(t *testing.T) {
	d := NewGetTrustStoreListDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_trust_store_list" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_trust_store_list")
	}
}
