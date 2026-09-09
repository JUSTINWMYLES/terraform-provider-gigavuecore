package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSnmpThrottleStatusDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadSnmpThrottleStatusDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadSnmpThrottleStatusDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadSnmpThrottleStatusDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadSnmpThrottleStatusDataSourceMetadata(t *testing.T) {
	d := NewLoadSnmpThrottleStatusDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_snmp_throttle_status" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_snmp_throttle_status")
	}
}
