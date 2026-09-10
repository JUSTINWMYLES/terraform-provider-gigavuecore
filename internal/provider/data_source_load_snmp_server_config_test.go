package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadSnmpServerConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadSnmpServerConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadSnmpServerConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadSnmpServerConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadSnmpServerConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadSnmpServerConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_snmp_server_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_snmp_server_config")
	}
}
