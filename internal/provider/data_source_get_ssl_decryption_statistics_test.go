package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSslDecryptionStatisticsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSslDecryptionStatisticsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSslDecryptionStatisticsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSslDecryptionStatisticsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSslDecryptionStatisticsDataSourceMetadata(t *testing.T) {
	d := NewGetSslDecryptionStatisticsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ssl_decryption_statistics" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ssl_decryption_statistics")
	}
}
