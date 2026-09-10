package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCommonCiphersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCommonCiphersDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCommonCiphersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCommonCiphersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCommonCiphersDataSourceMetadata(t *testing.T) {
	d := NewGetCommonCiphersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_common_ciphers" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_common_ciphers")
	}
}
