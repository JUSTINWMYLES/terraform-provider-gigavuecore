package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestRefreshAllClusterLicensesDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestRefreshAllClusterLicensesDataSourceSchemaValidation(t *testing.T) {
	d := NewRefreshAllClusterLicensesDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestRefreshAllClusterLicensesDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestRefreshAllClusterLicensesDataSourceMetadata(t *testing.T) {
	d := NewRefreshAllClusterLicensesDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_refresh_all_cluster_licenses" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_refresh_all_cluster_licenses")
	}
}
