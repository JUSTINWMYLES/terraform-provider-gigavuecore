package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadLicensingConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadLicensingConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadLicensingConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadLicensingConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadLicensingConfigDataSourceMetadata(t *testing.T) {
	d := NewLoadLicensingConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_licensing_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_licensing_config")
	}
}
