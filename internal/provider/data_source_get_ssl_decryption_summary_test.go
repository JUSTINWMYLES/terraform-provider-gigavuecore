package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslDecryptionSummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSslDecryptionSummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSslDecryptionSummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSslDecryptionSummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSslDecryptionSummaryDataSourceMetadata(t *testing.T) {
	d := NewGetSslDecryptionSummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ssl_decryption_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ssl_decryption_summary")
	}
}
