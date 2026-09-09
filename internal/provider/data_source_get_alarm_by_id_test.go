package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAlarmByIdDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAlarmByIdDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAlarmByIdDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAlarmByIdDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAlarmByIdDataSourceMetadata(t *testing.T) {
	d := NewGetAlarmByIdDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_alarm_by_id" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_alarm_by_id")
	}
}
