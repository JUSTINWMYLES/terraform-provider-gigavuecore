package provider

import (
	"context"
	"testing"
)
import tfframeworkdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"

// TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceSchemaValidation verifies that the generated data source schema is valid.
func TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceSchemaValidation(t *testing.T) {
	d := NewGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource()
	var resp tfframeworkdatasource.SchemaResponse
	d.Schema(context.Background(), tfframeworkdatasource.SchemaRequest{}, &resp)
	diags := resp.Schema.ValidateImplementation(context.Background())
	if diags.HasError() {
		t.Fatalf("schema validation failed: %s", diags)
	}
}

// TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceMetadata verifies that the generated data source reports the expected type name.
func TestGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSourceMetadata(t *testing.T) {
	d := NewGetConnectedLinksTrafficFlowsTroubleshootFlowsTopologyDataSource()
	var resp tfframeworkdatasource.MetadataResponse
	d.Metadata(context.Background(), tfframeworkdatasource.MetadataRequest{}, &resp)
	if resp.TypeName != "gigavuecore_get_connected_links_traffic_flows_troubleshoot_flows_topology" {
		t.Fatalf("TypeName = %q, want %q", resp.TypeName, "gigavuecore_get_connected_links_traffic_flows_troubleshoot_flows_topology")
	}
}
