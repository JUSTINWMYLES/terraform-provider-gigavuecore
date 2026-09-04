package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllSslDecryptionSummaryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllSslDecryptionSummaryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllSslDecryptionSummaryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllSslDecryptionSummaryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllSslDecryptionSummaryDataSourceMetadata(t *testing.T) {
	d := NewGetAllSslDecryptionSummaryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_ssl_decryption_summary" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_ssl_decryption_summary")
	}
}
