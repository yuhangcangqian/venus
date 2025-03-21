package gengen_test

import (
	"context"
	"os"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin"

	"github.com/filecoin-project/go-state-types/abi"
	blockstore "github.com/ipfs/boxo/blockstore"
	ds "github.com/ipfs/go-datastore"

	"github.com/filecoin-project/venus/pkg/constants"
	th "github.com/filecoin-project/venus/pkg/testhelpers"
	tf "github.com/filecoin-project/venus/pkg/testhelpers/testflags"
	genutil "github.com/filecoin-project/venus/tools/gengen/util"

	_ "github.com/filecoin-project/venus/pkg/crypto/bls"
	_ "github.com/filecoin-project/venus/pkg/crypto/delegated" // enable delegated signatures
	_ "github.com/filecoin-project/venus/pkg/crypto/secp"
	blockstoreutil "github.com/filecoin-project/venus/venus-shared/blockstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testConfig(t *testing.T) *genutil.GenesisCfg {
	fiftyCommCfgs, err := genutil.MakeCommitCfgs(50)
	require.NoError(t, err)
	tenCommCfgs, err := genutil.MakeCommitCfgs(10)
	require.NoError(t, err)

	return &genutil.GenesisCfg{
		KeysToGen:         4,
		PreallocatedFunds: []string{"1000000", "500000"},
		Miners: []*genutil.CreateStorageMinerConfig{
			{
				Owner:            0,
				CommittedSectors: fiftyCommCfgs,
				SealProofType:    constants.DevSealProofType,
				MarketBalance:    abi.NewTokenAmount(0),
			},
			{
				Owner:            1,
				CommittedSectors: tenCommCfgs,
				SealProofType:    constants.DevSealProofType,
				MarketBalance:    abi.NewTokenAmount(0),
			},
		},
		Network: "gfctest",
		Seed:    defaultSeed,
		Time:    defaultTime,
	}
}

const (
	defaultSeed = 4
	defaultTime = 123456789
)

func TestGenGenLoading(t *testing.T) {
	tf.IntegrationTest(t)
	address.CurrentNetwork = address.Testnet

	fi, err := os.CreateTemp("", "gengentest")
	assert.NoError(t, err)

	_, err = genutil.GenGenesisCar(testConfig(t), fi)
	assert.NoError(t, err)
	assert.NoError(t, fi.Close())

	td := th.NewDaemon(t, th.GenesisFile(fi.Name())).Start()
	defer td.ShutdownSuccess()

	o := td.Run("state", "list-actor").AssertSuccess()

	stdout := o.ReadStdout()
	//address won't change
	assert.Contains(t, stdout, builtin.StoragePowerActorAddr.String())
	assert.Contains(t, stdout, builtin.StorageMarketActorAddr.String())
	assert.Contains(t, stdout, builtin.InitActorAddr.String())
}

func TestGenGenDeterministic(t *testing.T) {
	tf.IntegrationTest(t)

	ctx := context.Background()
	var info *genutil.RenderedGenInfo
	for i := 0; i < 5; i++ {
		bstore := blockstore.NewBlockstore(ds.NewMapDatastore())
		inf, err := genutil.GenGen(ctx, testConfig(t), blockstoreutil.Adapt(bstore))
		assert.NoError(t, err)
		if info == nil {
			info = inf
		} else {
			assert.Equal(t, info.GenesisCid, inf.GenesisCid)
			assert.Equal(t, info.Miners, inf.Miners)
			assert.Equal(t, len(info.Keys), len(inf.Keys))
			for i, key := range inf.Keys {
				assert.Equal(t, info.Keys[i].Key(), key.Key())
			}
		}
	}
}
