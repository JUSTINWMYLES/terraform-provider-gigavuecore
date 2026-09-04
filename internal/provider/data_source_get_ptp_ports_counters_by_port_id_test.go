package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPtpPortsCountersByPortIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPtpPortsCountersByPortIdDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPtpPortsCountersByPortIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPtpPortsCountersByPortIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPtpPortsCountersByPortIdDataSourceMetadata(t *testing.T) {
	d := NewGetPtpPortsCountersByPortIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ptp_ports_counters_by_port_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ptp_ports_counters_by_port_id")
	}
}
