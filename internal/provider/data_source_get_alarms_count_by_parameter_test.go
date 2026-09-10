package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAlarmsCountByParameterDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAlarmsCountByParameterDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAlarmsCountByParameterDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAlarmsCountByParameterDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAlarmsCountByParameterDataSourceMetadata(t *testing.T) {
	d := NewGetAlarmsCountByParameterDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_alarms_count_by_parameter" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_alarms_count_by_parameter")
	}
}
