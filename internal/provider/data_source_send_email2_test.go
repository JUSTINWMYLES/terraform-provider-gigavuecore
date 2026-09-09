package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestSendEmail2DataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestSendEmail2DataSourceSchemaValidation(t *testing.T) {
	d := NewSendEmail2DataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestSendEmail2DataSourceMetadata verifies that the generated data source reports the expected type name.
func TestSendEmail2DataSourceMetadata(t *testing.T) {
	d := NewSendEmail2DataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_send_email2" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_send_email2")
	}
}
