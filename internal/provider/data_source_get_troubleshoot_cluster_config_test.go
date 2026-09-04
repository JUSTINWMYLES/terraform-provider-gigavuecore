package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetTroubleshootClusterConfigDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetTroubleshootClusterConfigDataSourceSchemaValidation(t *testing.T) {
	d := NewGetTroubleshootClusterConfigDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetTroubleshootClusterConfigDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetTroubleshootClusterConfigDataSourceMetadata(t *testing.T) {
	d := NewGetTroubleshootClusterConfigDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_troubleshoot_cluster_config" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_troubleshoot_cluster_config")
	}
}
