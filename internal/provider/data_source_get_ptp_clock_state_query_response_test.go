package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPtpClockStateQueryResponseDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPtpClockStateQueryResponseDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPtpClockStateQueryResponseDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPtpClockStateQueryResponseDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPtpClockStateQueryResponseDataSourceMetadata(t *testing.T) {
	d := NewGetPtpClockStateQueryResponseDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_ptp_clock_state_query_response" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_ptp_clock_state_query_response")
	}
}
