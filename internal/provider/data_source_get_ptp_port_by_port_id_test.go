package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPtpPortByPortIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPtpPortByPortIdDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPtpPortByPortIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPtpPortByPortIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPtpPortByPortIdDataSourceMetadata(t *testing.T) {
	d := NewGetPtpPortByPortIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ptp_port_by_port_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ptp_port_by_port_id")
	}
}
