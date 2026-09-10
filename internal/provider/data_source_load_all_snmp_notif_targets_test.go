package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadAllSnmpNotifTargetsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadAllSnmpNotifTargetsDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadAllSnmpNotifTargetsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadAllSnmpNotifTargetsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadAllSnmpNotifTargetsDataSourceMetadata(t *testing.T) {
	d := NewLoadAllSnmpNotifTargetsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_all_snmp_notif_targets" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_all_snmp_notif_targets")
	}
}
