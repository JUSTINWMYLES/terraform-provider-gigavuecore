package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestExpiryNotifAndEmailAndCountDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestExpiryNotifAndEmailAndCountDataSourceSchemaValidation(t *testing.T) {
	d := NewExpiryNotifAndEmailAndCountDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestExpiryNotifAndEmailAndCountDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestExpiryNotifAndEmailAndCountDataSourceMetadata(t *testing.T) {
	d := NewExpiryNotifAndEmailAndCountDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_expiry_notif_and_email_and_count" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_expiry_notif_and_email_and_count")
	}
}
