package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllPtpPortsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllPtpPortsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllPtpPortsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllPtpPortsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllPtpPortsDataSourceMetadata(t *testing.T) {
	d := NewGetAllPtpPortsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_ptp_ports" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_ptp_ports")
	}
}
