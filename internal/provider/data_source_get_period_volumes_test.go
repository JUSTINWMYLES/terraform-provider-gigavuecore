package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetPeriodVolumesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetPeriodVolumesDataSourceSchemaValidation(t *testing.T) {
	d := NewGetPeriodVolumesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetPeriodVolumesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetPeriodVolumesDataSourceMetadata(t *testing.T) {
	d := NewGetPeriodVolumesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_period_volumes" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_period_volumes")
	}
}
