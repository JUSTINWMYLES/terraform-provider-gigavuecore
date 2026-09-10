package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllTimeStampingPtpConfigsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllTimeStampingPtpConfigsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllTimeStampingPtpConfigsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllTimeStampingPtpConfigsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllTimeStampingPtpConfigsDataSourceMetadata(t *testing.T) {
	d := NewGetAllTimeStampingPtpConfigsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_time_stamping_ptp_configs" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_time_stamping_ptp_configs")
	}
}
