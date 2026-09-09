package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetAllClusterMapsOfAnInternalFabricMapDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetAllClusterMapsOfAnInternalFabricMapDataSourceSchemaValidation(t *testing.T) {
	d := NewGetAllClusterMapsOfAnInternalFabricMapDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetAllClusterMapsOfAnInternalFabricMapDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetAllClusterMapsOfAnInternalFabricMapDataSourceMetadata(t *testing.T) {
	d := NewGetAllClusterMapsOfAnInternalFabricMapDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_all_cluster_maps_of_an_internal_fabric_map")
	}
}
