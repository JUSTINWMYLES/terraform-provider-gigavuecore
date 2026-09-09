package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEventNotificationTemplateDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetEventNotificationTemplateDataSourceSchemaValidation(t *testing.T) {
	d := NewGetEventNotificationTemplateDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetEventNotificationTemplateDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetEventNotificationTemplateDataSourceMetadata(t *testing.T) {
	d := NewGetEventNotificationTemplateDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_event_notification_template" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_event_notification_template")
	}
}
