package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllSslDecryptionKeyMapsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllSslDecryptionKeyMapsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllSslDecryptionKeyMapsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllSslDecryptionKeyMapsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllSslDecryptionKeyMapsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllSslDecryptionKeyMapsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_ssl_decryption_key_maps" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_ssl_decryption_key_maps")
	}
}
