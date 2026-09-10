package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllPtpPortsCountersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllPtpPortsCountersDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllPtpPortsCountersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllPtpPortsCountersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllPtpPortsCountersDataSourceMetadata(t *testing.T) {
	d := NewGetAllPtpPortsCountersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_ptp_ports_counters" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_ptp_ports_counters")
	}
}
