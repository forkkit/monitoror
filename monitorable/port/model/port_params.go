//+build !faker

package model

type (
	PortParams struct {
		Hostname string `json:"hostname" query:"hostname"`
		Port     int    `json:"port" query:"port"`
	}
)

func (p *PortParams) Validate() bool {
	return p.Hostname != "" && p.Port != 0
}