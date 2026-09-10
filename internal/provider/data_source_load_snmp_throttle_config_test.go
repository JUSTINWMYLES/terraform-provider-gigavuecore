package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSnmpThrottleConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadSnmpThrottleConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadSnmpThrottleConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadSnmpThrottleConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadSnmpThrottleConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadSnmpThrottleConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_snmp_throttle_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_snmp_throttle_config")
	}
}
