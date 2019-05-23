package cbortests

import (
	"testing"

	multihash "github.com/multiformats/go-multihash"

	cbor "github.com/ipfs/go-ipld-cbor"
	"github.com/quorumcontrol/cboroneofpb/pb/oneoftest"
	"github.com/stretchr/testify/require"
)

func init() {
	cbor.RegisterCborType(oneoftest.DemoBrokenCbor{})
	cbor.RegisterCborType(oneoftest.OneKind{})
	cbor.RegisterCborType(oneoftest.DemoBrokenCbor_OneKind{})
	cbor.RegisterCborType(oneoftest.AnotherKind{})
	cbor.RegisterCborType(oneoftest.DemoBrokenCbor_AnotherKind{})
	cbor.RegisterCborType(struct{}{})
}

func TestRoundtrip(t *testing.T) {
	obj := &oneoftest.DemoBrokenCbor{
		Payload: &oneoftest.DemoBrokenCbor_OneKind{
			OneKind: &oneoftest.OneKind{
				Value: "hi",
			},
		},
	}

	n, err := cbor.WrapObject(obj, multihash.SHA2_256, -1)
	require.Nil(t, err)

	newObj := &oneoftest.DemoBrokenCbor{}
	err = cbor.DecodeInto(n.RawData(), newObj)
	require.Nil(t, err)

	require.Equal(t, obj, newObj)
}

// Outputs:

// 13:28 ~/code/cboroneofpb (master) $ go test ./cbortest
// --- FAIL: TestRoundtrip (0.00s)
// panic: reflect.Set: value of type map[string]interface {} is not assignable to type oneoftest.isDemoBrokenCbor_Payload [recovered]
// 	panic: reflect.Set: value of type map[string]interface {} is not assignable to type oneoftest.isDemoBrokenCbor_Payload

// goroutine 34 [running]:
// testing.tRunner.func1(0xc000146100)
// 	/usr/local/Cellar/go/1.12.4/libexec/src/testing/testing.go:830 +0x392
// panic(0x13e1c40, 0xc00013e3d0)
// 	/usr/local/Cellar/go/1.12.4/libexec/src/runtime/panic.go:522 +0x1b5
// reflect.Value.assignTo(0x13fc0e0, 0xc000142390, 0x15, 0x1473a13, 0xb, 0x140de00, 0xc000142300, 0x100edfd, 0x1457fe0, 0xc000142390)
// 	/usr/local/Cellar/go/1.12.4/libexec/src/reflect/value.go:2339 +0x437
// reflect.Value.Set(0x140de00, 0xc000142300, 0x194, 0x13fc0e0, 0xc000142390, 0x15)
// 	/usr/local/Cellar/go/1.12.4/libexec/src/reflect/value.go:1473 +0xa8
// github.com/polydawn/refmt/obj.(*unmarshalMachineWildcard).prepareDemux(0xc00017a890, 0xc000158180, 0xc000158180, 0xc00016e1c0, 0x17d5a40, 0x1e01460, 0xc000162bf8)
// 	/Users/tobowers/code/go/pkg/mod/github.com/polydawn/refmt@v0.0.0-20190221155625-df39d6c2d992/obj/unmarshalWildcard.go:65 +0x5c0
// github.com/polydawn/refmt/obj.(*unmarshalMachineWildcard).Step(0xc00017a890, 0xc000158180, 0xc000158180, 0xc00016e1c0, 0xc000162c28, 0x11c3656, 0xc00017a838)
// 	/Users/tobowers/code/go/pkg/mod/github.com/polydawn/refmt@v0.0.0-20190221155625-df39d6c2d992/obj/unmarshalWildcard.go:27 +0x144
// github.com/polydawn/refmt/obj.(*Unmarshaller).Step(0xc000158180, 0xc00016e1c0, 0x140de00, 0xc000142300, 0x194)
// 	/Users/tobowers/code/go/pkg/mod/github.com/polydawn/refmt@v0.0.0-20190221155625-df39d6c2d992/obj/unmarshal.go:60 +0x4c
// github.com/polydawn/refmt/obj.(*Unmarshaller).Recurse(0xc000158180, 0xc00016e1c0, 0x140de00, 0xc000142300, 0x194, 0x1507d00, 0x140de00, 0x14fc480, 0xc00017a890, 0x1, ...)
// 	/Users/tobowers/code/go/pkg/mod/github.com/polydawn/refmt@v0.0.0-20190221155625-df39d6c2d992/obj/unmarshal.go:106 +0x116
// github.com/polydawn/refmt/obj.(*unmarshalMachineStructAtlas).Step(0xc00017a6f0, 0xc000158180, 0xc000158180, 0xc00016e1c0, 0x145d400, 0x0, 0x0)
// 	/Users/tobowers/code/go/pkg/mod/github.com/polydawn/refmt@v0.0.0-20190221155625-df39d6c2d992/obj/unmarshalStruct.go:66 +0x62b
// github.com/polydawn/refmt/obj.(*Unmarshaller).Step(0xc000158180, 0xc00016e1c0, 0x100e100, 0x0, 0x0)
// 	/Users/tobowers/code/go/pkg/mod/github.com/polydawn/refmt@v0.0.0-20190221155625-df39d6c2d992/obj/unmarshal.go:60 +0x4c
// github.com/polydawn/refmt/shared.TokenPump.Run(0x14f8760, 0xc0001700a0, 0x14f87e0, 0xc000158180, 0x0, 0xc000142360)
// 	/Users/tobowers/code/go/pkg/mod/github.com/polydawn/refmt@v0.0.0-20190221155625-df39d6c2d992/shared/pump.go:35 +0x9b
// github.com/polydawn/refmt/cbor.(*Unmarshaller).Unmarshal(0xc000142330, 0x1451f60, 0xc000142300, 0x0, 0xc00015cd80)
// 	/Users/tobowers/code/go/pkg/mod/github.com/polydawn/refmt@v0.0.0-20190221155625-df39d6c2d992/cbor/cborHelpers.go:90 +0x85
// github.com/ipfs/go-ipld-cbor/encoding.(*Unmarshaller).Decode(...)
// 	/Users/tobowers/code/go/pkg/mod/github.com/ipfs/go-ipld-cbor@v0.0.2/encoding/unmarshaller.go:37
// github.com/ipfs/go-ipld-cbor/encoding.(*Unmarshaller).Unmarshal(0xc000140080, 0xc000176000, 0x8a, 0x94, 0x1451f60, 0xc000142300, 0x100e1e8, 0x30)
// 	/Users/tobowers/code/go/pkg/mod/github.com/ipfs/go-ipld-cbor@v0.0.2/encoding/unmarshaller.go:44 +0xb5
// github.com/ipfs/go-ipld-cbor/encoding.(*PooledUnmarshaller).Unmarshal(0x17d43b0, 0xc000176000, 0x8a, 0x94, 0x1451f60, 0xc000142300, 0x0, 0x5ce683ff)
// 	/Users/tobowers/code/go/pkg/mod/github.com/ipfs/go-ipld-cbor@v0.0.2/encoding/unmarshaller.go:76 +0x94
// github.com/ipfs/go-ipld-cbor.DecodeInto(...)
// 	/Users/tobowers/code/go/pkg/mod/github.com/ipfs/go-ipld-cbor@v0.0.2/node.go:108
// github.com/quorumcontrol/cboroneofpb/cbortest.TestRoundtrip(0xc000146100)
// 	/Users/tobowers/code/cboroneofpb/cbortest/cbortest_test.go:35 +0x187
// testing.tRunner(0xc000146100, 0x148b390)
// 	/usr/local/Cellar/go/1.12.4/libexec/src/testing/testing.go:865 +0xc0
// created by testing.(*T).Run
// 	/usr/local/Cellar/go/1.12.4/libexec/src/testing/testing.go:916 +0x35a
// FAIL	github.com/quorumcontrol/cboroneofpb/cbortest	0.016s
