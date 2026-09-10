package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAlarmSuppressionMetadataDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllAlarmSuppressionMetadataDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllAlarmSuppressionMetadataDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllAlarmSuppressionMetadataDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllAlarmSuppressionMetadataDataSourceMetadata(t *testing.T) {
	d := NewGetAllAlarmSuppressionMetadataDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_alarm_suppression_metadata" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_alarm_suppression_metadata")
	}
}
