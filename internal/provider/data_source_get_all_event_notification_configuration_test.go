package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllEventNotificationConfigurationDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllEventNotificationConfigurationDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllEventNotificationConfigurationDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllEventNotificationConfigurationDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllEventNotificationConfigurationDataSourceMetadata(t *testing.T) {
	d := NewGetAllEventNotificationConfigurationDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_event_notification_configuration" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_event_notification_configuration")
	}
}
