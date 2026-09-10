package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAcmeCertificateDetailsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAcmeCertificateDetailsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAcmeCertificateDetailsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAcmeCertificateDetailsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAcmeCertificateDetailsDataSourceMetadata(t *testing.T) {
	d := NewGetAcmeCertificateDetailsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_acme_certificate_details" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_acme_certificate_details")
	}
}
