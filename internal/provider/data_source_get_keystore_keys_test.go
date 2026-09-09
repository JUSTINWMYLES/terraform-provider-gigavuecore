package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetKeystoreKeysDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetKeystoreKeysDataSourceSchemaValidation(t *testing.T) {
	d := NewGetKeystoreKeysDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetKeystoreKeysDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetKeystoreKeysDataSourceMetadata(t *testing.T) {
	d := NewGetKeystoreKeysDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_keystore_keys" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_keystore_keys")
	}
}
