package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllNtpServerDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllNtpServerDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllNtpServerDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllNtpServerDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllNtpServerDataSourceMetadata(t *testing.T) {
	d := NewGetAllNtpServerDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_ntp_server" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_ntp_server")
	}
}
