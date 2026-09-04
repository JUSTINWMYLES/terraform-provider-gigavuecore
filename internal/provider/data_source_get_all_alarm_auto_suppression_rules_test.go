package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllAlarmAutoSuppressionRulesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllAlarmAutoSuppressionRulesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllAlarmAutoSuppressionRulesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllAlarmAutoSuppressionRulesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllAlarmAutoSuppressionRulesDataSourceMetadata(t *testing.T) {
	d := NewGetAllAlarmAutoSuppressionRulesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_alarm_auto_suppression_rules" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_alarm_auto_suppression_rules")
	}
}
