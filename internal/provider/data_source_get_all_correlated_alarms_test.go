package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllCorrelatedAlarmsDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllCorrelatedAlarmsDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllCorrelatedAlarmsDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllCorrelatedAlarmsDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllCorrelatedAlarmsDataSourceMetadata(t *testing.T) {
	d := NewGetAllCorrelatedAlarmsDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_correlated_alarms" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_correlated_alarms")
	}
}
