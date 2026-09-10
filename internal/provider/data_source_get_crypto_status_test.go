package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCryptoStatusDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCryptoStatusDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCryptoStatusDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCryptoStatusDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCryptoStatusDataSourceMetadata(t *testing.T) {
	d := NewGetCryptoStatusDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_crypto_status" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_crypto_status")
	}
}
