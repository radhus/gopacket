package layers

import (
	"encoding/binary"
	"github.com/google/gopacket"
)

type LCPCode uint8

const (
	LCPCodeConfigureRequest LCPCode = 1
	LCPCodeConfigureAck     LCPCode = 2
	LCPCodeConfigureNak     LCPCode = 3
	LCPCodeConfigureReject  LCPCode = 4
	LCPCodeTerminateRequest LCPCode = 5
	LCPCodeTerminateAck     LCPCode = 6
	LCPCodeCodeReject       LCPCode = 7
	LCPCodeProtocolReject   LCPCode = 8
	LCPCodeEchoRequest      LCPCode = 9
	LCPCodeEchoReply        LCPCode = 10
	LCPCodeDiscardRequest   LCPCode = 11
	LCPCodeIdentification   LCPCode = 12
	LCPCodeTimeRemaining    LCPCode = 13
)

func (c LCPCode) String() string {
	switch c {
	case LCPCodeConfigureRequest:
		return "Configure-Request"
	case LCPCodeConfigureAck    :
		return "Configure-Ack"
	case LCPCodeConfigureNak    :
		return "Configure-Nak"
	case LCPCodeConfigureReject :
		return "Configure-Reject"
	case LCPCodeTerminateRequest:
		return "Terminate-Request"
	case LCPCodeTerminateAck    :
		return "Terminate-Ack"
	case LCPCodeCodeReject      :
		return "Code-Reject"
	case LCPCodeProtocolReject  :
		return "Protocol-Reject"
	case LCPCodeEchoRequest     :
		return "Echo-Request"
	case LCPCodeEchoReply       :
		return "Echo-Reply"
	case LCPCodeDiscardRequest  :
		return "Discard-Request"
	case LCPCodeIdentification  :
		return "Identification"
	case LCPCodeTimeRemaining   :
		return "Time-Remaining"
	default:
		return "Unknown"
	}
}

type LCP struct {
	BaseLayer
	Code LCPCode
	Identifier uint8
	Length uint16
}

func (l *LCP) LayerType() gopacket.LayerType { return LayerTypeLCP }

func (l *LCP) Payload() []byte {
	return l.BaseLayer.Payload
}

func (l *LCP) DecodeFromBytes(data []byte, df gopacket.DecodeFeedback) error {
	l.Code = LCPCode(data[0])
	l.Identifier = data[1]
	l.Length = binary.BigEndian.Uint16(data[2:4])
	l.BaseLayer.Contents = data[:4]
	l.BaseLayer.Payload = data[4:]
	return nil
}

func decodeLCP(data []byte, p gopacket.PacketBuilder) error {
	lcp := &LCP{}
	err := lcp.DecodeFromBytes(data, p)
	if err != nil {
		return err
	}
	p.AddLayer(lcp)
	p.SetApplicationLayer(lcp)
	return nil
}

func (l *LCP) SerializeTo(b gopacket.SerializeBuffer, opts gopacket.SerializeOptions) error {
	bytes, err := b.PrependBytes(4)
	if err != nil {
		return err
	}

	bytes[0] = uint8(l.Code)
	bytes[1] = l.Identifier
	if (opts.FixLengths) {
		l.Length = uint16(len(b.Bytes()))
	}
	binary.BigEndian.PutUint16(bytes[2:], l.Length)
	return nil
}