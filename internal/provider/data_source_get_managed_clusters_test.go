package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetManagedClustersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetManagedClustersDataSourceSchemaValidation(t *testing.T) {
	d := NewGetManagedClustersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetManagedClustersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetManagedClustersDataSourceMetadata(t *testing.T) {
	d := NewGetManagedClustersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_managed_clusters" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_managed_clusters")
	}
}
