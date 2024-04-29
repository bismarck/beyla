// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package ebpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type NetSkFlowId NetSkFlowIdT

type NetSkFlowIdT struct {
	SrcIp             struct{ In6U struct{ U6Addr8 [16]uint8 } }
	DstIp             struct{ In6U struct{ U6Addr8 [16]uint8 } }
	EthProtocol       uint16
	Direction         uint8
	SrcPort           uint16
	DstPort           uint16
	TransportProtocol uint8
	IfIndex           uint32
}

type NetSkFlowMetrics NetSkFlowMetricsT

type NetSkFlowMetricsT struct {
	Packets         uint32
	Bytes           uint64
	StartMonoTimeNs uint64
	EndMonoTimeNs   uint64
	Flags           uint16
	Errno           uint8
}

type NetSkFlowRecordT struct {
	Id      NetSkFlowId
	Metrics NetSkFlowMetrics
}

// LoadNetSk returns the embedded CollectionSpec for NetSk.
func LoadNetSk() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_NetSkBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load NetSk: %w", err)
	}

	return spec, err
}

// LoadNetSkObjects loads NetSk and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*NetSkObjects
//	*NetSkPrograms
//	*NetSkMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadNetSkObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadNetSk()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// NetSkSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type NetSkSpecs struct {
	NetSkProgramSpecs
	NetSkMapSpecs
}

// NetSkSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type NetSkProgramSpecs struct {
	SocketHttpFilter *ebpf.ProgramSpec `ebpf:"socket__http_filter"`
}

// NetSkMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type NetSkMapSpecs struct {
	AggregatedFlows *ebpf.MapSpec `ebpf:"aggregated_flows"`
	DirectFlows     *ebpf.MapSpec `ebpf:"direct_flows"`
}

// NetSkObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadNetSkObjects or ebpf.CollectionSpec.LoadAndAssign.
type NetSkObjects struct {
	NetSkPrograms
	NetSkMaps
}

func (o *NetSkObjects) Close() error {
	return _NetSkClose(
		&o.NetSkPrograms,
		&o.NetSkMaps,
	)
}

// NetSkMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadNetSkObjects or ebpf.CollectionSpec.LoadAndAssign.
type NetSkMaps struct {
	AggregatedFlows *ebpf.Map `ebpf:"aggregated_flows"`
	DirectFlows     *ebpf.Map `ebpf:"direct_flows"`
}

func (m *NetSkMaps) Close() error {
	return _NetSkClose(
		m.AggregatedFlows,
		m.DirectFlows,
	)
}

// NetSkPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadNetSkObjects or ebpf.CollectionSpec.LoadAndAssign.
type NetSkPrograms struct {
	SocketHttpFilter *ebpf.Program `ebpf:"socket__http_filter"`
}

func (p *NetSkPrograms) Close() error {
	return _NetSkClose(
		p.SocketHttpFilter,
	)
}

func _NetSkClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed netsk_bpfel_arm64.o
var _NetSkBytes []byte
