package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllPtpCountersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllPtpCountersDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllPtpCountersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllPtpCountersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllPtpCountersDataSourceMetadata(t *testing.T) {
	d := NewGetAllPtpCountersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_ptp_counters" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_ptp_counters")
	}
}
