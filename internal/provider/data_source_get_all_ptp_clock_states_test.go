package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllPtpClockStatesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllPtpClockStatesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllPtpClockStatesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllPtpClockStatesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllPtpClockStatesDataSourceMetadata(t *testing.T) {
	d := NewGetAllPtpClockStatesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_ptp_clock_states" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_ptp_clock_states")
	}
}
