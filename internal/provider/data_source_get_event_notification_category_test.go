package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEventNotificationCategoryDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetEventNotificationCategoryDataSourceSchemaValidation(t *testing.T) {
	d := NewGetEventNotificationCategoryDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetEventNotificationCategoryDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetEventNotificationCategoryDataSourceMetadata(t *testing.T) {
	d := NewGetEventNotificationCategoryDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_event_notification_category" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_event_notification_category")
	}
}
