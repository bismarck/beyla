// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package goruntime

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpfGoroutineMetadata struct {
	Parent    uint64
	Timestamp uint64
}

type bpfHttpConnectionMetadataT struct {
	Pid struct {
		HostPid   uint32
		UserPid   uint32
		Namespace uint32
	}
	Type uint8
}

type bpfNewFuncInvocationT struct{ Parent uint64 }

type bpfPidConnectionInfoT struct {
	Conn bpfConnectionInfoT
	Pid  uint32
}

type bpfPidKeyT struct {
	Pid       uint32
	Namespace uint32
}

type bpfTpInfoPidT struct {
	Tp    bpfTpInfoT
	Pid   uint32
	Valid uint8
	_     [3]byte
}

type bpfTpInfoT struct {
	TraceId  [16]uint8
	SpanId   [8]uint8
	ParentId [8]uint8
	Ts       uint64
	Flags    uint8
	_        [7]byte
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	UprobeProcGoexit1     *ebpf.ProgramSpec `ebpf:"uprobe_proc_goexit1"`
	UprobeProcNewproc1    *ebpf.ProgramSpec `ebpf:"uprobe_proc_newproc1"`
	UprobeProcNewproc1Ret *ebpf.ProgramSpec `ebpf:"uprobe_proc_newproc1_ret"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	Events                       *ebpf.MapSpec `ebpf:"events"`
	FilteredConnections          *ebpf.MapSpec `ebpf:"filtered_connections"`
	GoTraceMap                   *ebpf.MapSpec `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap    *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	Newproc1                     *ebpf.MapSpec `ebpf:"newproc1"`
	OngoingGoroutines            *ebpf.MapSpec `ebpf:"ongoing_goroutines"`
	OngoingHttpServerConnections *ebpf.MapSpec `ebpf:"ongoing_http_server_connections"`
	PidCache                     *ebpf.MapSpec `ebpf:"pid_cache"`
	TraceMap                     *ebpf.MapSpec `ebpf:"trace_map"`
	ValidPids                    *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	Events                       *ebpf.Map `ebpf:"events"`
	FilteredConnections          *ebpf.Map `ebpf:"filtered_connections"`
	GoTraceMap                   *ebpf.Map `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap    *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	Newproc1                     *ebpf.Map `ebpf:"newproc1"`
	OngoingGoroutines            *ebpf.Map `ebpf:"ongoing_goroutines"`
	OngoingHttpServerConnections *ebpf.Map `ebpf:"ongoing_http_server_connections"`
	PidCache                     *ebpf.Map `ebpf:"pid_cache"`
	TraceMap                     *ebpf.Map `ebpf:"trace_map"`
	ValidPids                    *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.Events,
		m.FilteredConnections,
		m.GoTraceMap,
		m.GolangMapbucketStorageMap,
		m.Newproc1,
		m.OngoingGoroutines,
		m.OngoingHttpServerConnections,
		m.PidCache,
		m.TraceMap,
		m.ValidPids,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	UprobeProcGoexit1     *ebpf.Program `ebpf:"uprobe_proc_goexit1"`
	UprobeProcNewproc1    *ebpf.Program `ebpf:"uprobe_proc_newproc1"`
	UprobeProcNewproc1Ret *ebpf.Program `ebpf:"uprobe_proc_newproc1_ret"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.UprobeProcGoexit1,
		p.UprobeProcNewproc1,
		p.UprobeProcNewproc1Ret,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_x86.o
var _BpfBytes []byte
