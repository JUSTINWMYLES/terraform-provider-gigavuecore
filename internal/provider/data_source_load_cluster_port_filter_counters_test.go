package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestLoadClusterPortFilterCountersDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestLoadClusterPortFilterCountersDataSourceSchemaValidation(t *testing.T) {
	d := NewLoadClusterPortFilterCountersDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestLoadClusterPortFilterCountersDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestLoadClusterPortFilterCountersDataSourceMetadata(t *testing.T) {
	d := NewLoadClusterPortFilterCountersDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_load_cluster_port_filter_counters" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_load_cluster_port_filter_counters")
	}
}
