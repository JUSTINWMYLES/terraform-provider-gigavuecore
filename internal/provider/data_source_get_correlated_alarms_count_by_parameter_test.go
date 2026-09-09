package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetCorrelatedAlarmsCountByParameterDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetCorrelatedAlarmsCountByParameterDataSourceSchemaValidation(t *testing.T) {
	d := NewGetCorrelatedAlarmsCountByParameterDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetCorrelatedAlarmsCountByParameterDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetCorrelatedAlarmsCountByParameterDataSourceMetadata(t *testing.T) {
	d := NewGetCorrelatedAlarmsCountByParameterDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_correlated_alarms_count_by_parameter" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_correlated_alarms_count_by_parameter")
	}
}
