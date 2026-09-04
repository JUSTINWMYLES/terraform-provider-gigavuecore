package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetKeystoreKeyCertDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetKeystoreKeyCertDataSourceSchemaValidation(t *testing.T) {
	d := NewGetKeystoreKeyCertDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetKeystoreKeyCertDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetKeystoreKeyCertDataSourceMetadata(t *testing.T) {
	d := NewGetKeystoreKeyCertDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_keystore_key_cert" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_keystore_key_cert")
	}
}
