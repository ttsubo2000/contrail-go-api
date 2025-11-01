//
// Automatically generated. DO NOT EDIT.
//

package types

type LacpParams struct {
	LacpEnable bool `json:"lacp_enable,omitempty"`
	LacpInterval string `json:"lacp_interval,omitempty"`
	LacpMode string `json:"lacp_mode,omitempty"`
}

type PortProfileParameters struct {
	PortParams *PortParameters `json:"port_params,omitempty"`
	FlowControl bool `json:"flow_control,omitempty"`
	LacpParams *LacpParams `json:"lacp_params,omitempty"`
	BpduLoopProtection bool `json:"bpdu_loop_protection,omitempty"`
	PortCosUntrust bool `json:"port_cos_untrust,omitempty"`
}
