package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSourceSchemaValidation(t *testing.T) {
	d := NewGetSerialNumbersOfChassisRequiringGvosLicenseDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetSerialNumbersOfChassisRequiringGvosLicenseDataSourceMetadata(t *testing.T) {
	d := NewGetSerialNumbersOfChassisRequiringGvosLicenseDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_serial_numbers_of_chassis_requiring_gvos_license" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_serial_numbers_of_chassis_requiring_gvos_license")
	}
}
