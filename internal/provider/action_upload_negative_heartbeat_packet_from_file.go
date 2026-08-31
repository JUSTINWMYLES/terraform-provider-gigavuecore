package provider

import "context"
import (
	action "github.com/hashicorp/terraform-plugin-framework/action"
	schema "github.com/hashicorp/terraform-plugin-framework/action/schema"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// Compile-time interface assertion for action.Action.
var _ action.Action = (*UploadNegativeHeartbeatPacketFromFileAction)(nil)

// UploadNegativeHeartbeatPacketFromFileAction is the generated Terraform action implementation.
type UploadNegativeHeartbeatPacketFromFileAction struct {
}

// UploadNegativeHeartbeatPacketFromFileActionModel describes the action configuration shape.
type UploadNegativeHeartbeatPacketFromFileActionModel struct {
	Alias     types.String `tfsdk:"alias"`
	ClusterId types.String `tfsdk:"cluster_id"`
	NhbPacket types.String `tfsdk:"nhb_packet" json:"nhbPacket"`
}

// NewUploadNegativeHeartbeatPacketFromFileAction returns a new instance of the generated action.
func NewUploadNegativeHeartbeatPacketFromFileAction() action.Action {
	return &UploadNegativeHeartbeatPacketFromFileAction{}
}

// Metadata returns the action type name.
func (r *UploadNegativeHeartbeatPacketFromFileAction) Metadata(_ context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "gigavuecore_upload_negative_heartbeat_packet_from_file"
}

// Schema returns the action schema.
func (r *UploadNegativeHeartbeatPacketFromFileAction) Schema(_ context.Context, _ action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{Description: "Upload a Negative Heartbeat Packet from local file", Attributes: map[string]schema.Attribute{"alias": schema.StringAttribute{MarkdownDescription: "alias of the target Negative Heartbeat Profile", Required: true}, "cluster_id": schema.StringAttribute{MarkdownDescription: "Target Cluster ID", Required: true}, "nhb_packet": schema.StringAttribute{MarkdownDescription: "Attached Negative Heartbeat Packet. In native binary format", Required: true}}}
}

// Invoke executes the action against the remote API.
// The generated Invoke method is intentionally stubbed; the remote API is not wired.
func (r *UploadNegativeHeartbeatPacketFromFileAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config UploadNegativeHeartbeatPacketFromFileActionModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.AddError("Generated provider scaffold", "Invoke is not wired to a remote API endpoint.")
}
