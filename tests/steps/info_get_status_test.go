package steps

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/make-software/casper-go-sdk/casper"
	"github.com/make-software/casper-go-sdk/rpc"
	"github.com/stormeye2000/cspr-sdk-standard-tests-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// The test features implementation for the info_get_peers.feature
func TestFeaturesInfoGetStatus(t *testing.T) {
	TestFeatures(t, "info_get_status.feature", InitializeInfoGetStatus)
}

func InitializeInfoGetStatus(ctx *godog.ScenarioContext) {

	var sdk casper.RPCClient
	var infoGetStatusResult rpc.InfoGetStatusResult
	var nctlGetStatus = casper.InfoGetStatusResult{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		utils.ReadConfig()
		sdk = utils.GetSdk()
		return ctx, nil
	})

	ctx.Step(`^that the info_get_status is invoked against nctl$`, func() error {

		err := utils.Pass

		infoGetStatusResult, err = sdk.GetStatus(context.Background())

		nctlGetStatus, err = utils.GetNodeStatus(1)

		return err
	})

	ctx.Step(`^an info_get_status_result is returned`, func() error {

		assert.NotNil(CasperT, infoGetStatusResult, "infoGetStatusResult is nil")

		return utils.Pass
	})

	ctx.Step(`^the info_get_status_result api_version is "([^"]*)"$`, func(apiVersion string) error {

		if apiVersion != infoGetStatusResult.APIVersion {
			return fmt.Errorf("expected %s ApiVersion to be %s", infoGetStatusResult.APIVersion, apiVersion)
		}
		return utils.Pass
	})

	ctx.Step(`^the info_get_status_result chainspec_name is "([^"]*)"$`, func(chainSpecName string) error {

		if chainSpecName != infoGetStatusResult.ChainSpecName {
			return fmt.Errorf("expected %s ChainSpecName to be %s", infoGetStatusResult.ChainSpecName, chainSpecName)
		}

		return utils.Pass
	})

	ctx.Step(`^the info_get_status_result has a valid last_added_block_info$`, func() error {

		err := expectEqual("hash", infoGetStatusResult.LastAddedBlockInfo.Hash, nctlGetStatus.LastAddedBlockInfo.Hash)

		if err == nil {
			err = expectEqual("Timestamp", infoGetStatusResult.LastAddedBlockInfo.Timestamp, nctlGetStatus.LastAddedBlockInfo.Timestamp)
		}

		if err == nil {
			err = expectEqual("EraID", infoGetStatusResult.LastAddedBlockInfo.EraID, nctlGetStatus.LastAddedBlockInfo.EraID)
		}

		if err == nil {
			err = expectEqual("Height", infoGetStatusResult.LastAddedBlockInfo.Height, nctlGetStatus.LastAddedBlockInfo.Height)
		}

		if err == nil {
			err = expectEqual("StateRootHash", infoGetStatusResult.LastAddedBlockInfo.StateRootHash, nctlGetStatus.LastAddedBlockInfo.StateRootHash)
		}

		if err == nil {
			err = expectEqual("Creator", infoGetStatusResult.LastAddedBlockInfo.Creator, nctlGetStatus.LastAddedBlockInfo.Creator)
		}

		return err
	})

	ctx.Step(`^the info_get_status_result has a valid our_public_signing_key$`, func() error {
		return expectEqual("our_public_signing_key", infoGetStatusResult.OutPublicSigningKey, nctlGetStatus.OutPublicSigningKey)
	})

	ctx.Step(`^the info_get_status_result has a valid starting_state_root_hash$`, func() error {
		return expectEqual("StartingStateRootHash", infoGetStatusResult.StartingStateRootHash, nctlGetStatus.StartingStateRootHash)
	})

	ctx.Step(`^the info_get_status_result has a valid build_version$`, func() error {
		return expectEqual("BuildVersion", infoGetStatusResult.BuildVersion, nctlGetStatus.BuildVersion)
	})

	ctx.Step(`^the info_get_status_result has a valid round_length$`, func() error {
		return expectEqual("RoundLength", infoGetStatusResult.RoundLength, nctlGetStatus.RoundLength)
	})

	ctx.Step(`^the info_get_status_result has a valid uptime$`, func() error {

		if !strings.HasSuffix(infoGetStatusResult.Uptime, "ms") {
			return fmt.Errorf("missing ms from uptime %s", infoGetStatusResult.Uptime)
		}

		return utils.Pass
	})

	ctx.Step(`^the info_get_status_result has a valid peers$`, func() error {
		return expectEqual("Peers", infoGetStatusResult.Peers, nctlGetStatus.Peers)
	})
}

func expectEqual(attribute string, actual any, expected any) error {

	if !assert.Equal(CasperT, expected, actual) {
		return fmt.Errorf("%s expected %s to be %s", attribute, actual, expected)
	}
	return utils.Pass
}