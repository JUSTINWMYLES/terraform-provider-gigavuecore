package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSerialNumberOfCardDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSerialNumberOfCardDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSerialNumberOfCardDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSerialNumberOfCardDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSerialNumberOfCardDataSourceMetadata(t *testing.T) {
	d := NewGetSerialNumberOfCardDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_serial_number_of_card" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_serial_number_of_card")
	}
}
