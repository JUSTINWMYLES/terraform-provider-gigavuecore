package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllClusterLicensesLicensingModuleAllFlatDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllClusterLicensesLicensingModuleAllFlatDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllClusterLicensesLicensingModuleAllFlatDataSourceMetadata(t *testing.T) {
	d := NewGetAllClusterLicensesLicensingModuleAllFlatDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_cluster_licenses_licensing_module_all_flat" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_cluster_licenses_licensing_module_all_flat")
	}
}
