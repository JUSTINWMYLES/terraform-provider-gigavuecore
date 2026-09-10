package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSupportedSnmpV3UserProtocolsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSupportedSnmpV3UserProtocolsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSupportedSnmpV3UserProtocolsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSupportedSnmpV3UserProtocolsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSupportedSnmpV3UserProtocolsDataSourceMetadata(t *testing.T) {
	d := NewGetSupportedSnmpV3UserProtocolsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_supported_snmp_v3_user_protocols" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_supported_snmp_v3_user_protocols")
	}
}
