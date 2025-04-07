// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshItemMtlsBackends struct {
	Conf      *MeshItemMtlsConf `tfsdk:"conf"`
	DpCert    *DpCert           `tfsdk:"dp_cert"`
	Mode      *Mode             `tfsdk:"mode"`
	Name      types.String      `tfsdk:"name"`
	RootChain *RootChain        `tfsdk:"root_chain"`
	Type      types.String      `tfsdk:"type"`
}
