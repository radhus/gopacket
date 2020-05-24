package layers

import "github.com/google/gopacket"

type IPCP struct {
	LCP
}

func (i *IPCP) LayerType() gopacket.LayerType { return LayerTypeIPCP }

func decodeIPCP(data []byte, p gopacket.PacketBuilder) error {
	ipcp := &IPCP{}
	err := ipcp.LCP.DecodeFromBytes(data, p)
	if err != nil {
		return err
	}
	p.AddLayer(ipcp)
	p.SetApplicationLayer(ipcp)
	return nil
}
