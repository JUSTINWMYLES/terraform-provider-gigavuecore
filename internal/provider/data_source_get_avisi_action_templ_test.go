package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAvisiActionTemplDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAvisiActionTemplDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAvisiActionTemplDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAvisiActionTemplDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAvisiActionTemplDataSourceMetadata(t *testing.T) {
	d := NewGetAvisiActionTemplDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_avisi_action_templ" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_avisi_action_templ")
	}
}
