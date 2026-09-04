package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAppIntelTemplateDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAppIntelTemplateDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAppIntelTemplateDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAppIntelTemplateDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAppIntelTemplateDataSourceMetadata(t *testing.T) {
	d := NewGetAppIntelTemplateDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_app_intel_template" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_app_intel_template")
	}
}
