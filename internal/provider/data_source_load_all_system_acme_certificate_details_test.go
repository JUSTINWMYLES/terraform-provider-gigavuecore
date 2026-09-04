package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllSystemAcmeCertificateDetailsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllSystemAcmeCertificateDetailsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllSystemAcmeCertificateDetailsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllSystemAcmeCertificateDetailsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllSystemAcmeCertificateDetailsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllSystemAcmeCertificateDetailsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_system_acme_certificate_details" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_system_acme_certificate_details")
	}
}
