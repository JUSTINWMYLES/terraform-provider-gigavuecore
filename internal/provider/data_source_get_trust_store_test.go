package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTrustStoreDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetTrustStoreDataSourceSchemaValidation(t *testing.T) {
	d := NewGetTrustStoreDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetTrustStoreDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetTrustStoreDataSourceMetadata(t *testing.T) {
	d := NewGetTrustStoreDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_trust_store" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_trust_store")
	}
}
