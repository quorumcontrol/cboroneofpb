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
