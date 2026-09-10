package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSystemDiagPsuPerSlotDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSystemDiagPsuPerSlotDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSystemDiagPsuPerSlotDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSystemDiagPsuPerSlotDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSystemDiagPsuPerSlotDataSourceMetadata(t *testing.T) {
	d := NewGetSystemDiagPsuPerSlotDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_system_diag_psu_per_slot" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_system_diag_psu_per_slot")
	}
}
