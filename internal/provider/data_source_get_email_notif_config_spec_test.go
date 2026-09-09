package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetEmailNotifConfigSpecDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetEmailNotifConfigSpecDataSourceSchemaValidation(t *testing.T) {
	d := NewGetEmailNotifConfigSpecDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetEmailNotifConfigSpecDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetEmailNotifConfigSpecDataSourceMetadata(t *testing.T) {
	d := NewGetEmailNotifConfigSpecDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_email_notif_config_spec" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_email_notif_config_spec")
	}
}
