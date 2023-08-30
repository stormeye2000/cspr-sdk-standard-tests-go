package steps

import (
	"context"
	"github.com/cucumber/godog"
	"github.com/make-software/casper-go-sdk/casper"
	"github.com/stormeye2000/cspr-sdk-standard-tests-go/tests/utils"
	"testing"
)

// The test features implementation for the info_get_peers.feature
func TestFeaturesDeploys(t *testing.T) {
	TestFeatures(t, "deploys.feature", InitializeDeploys)
}

func InitializeDeploys(ctx *godog.ScenarioContext) {

	var sdk casper.RPCClient

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		utils.ReadConfig()
		sdk = utils.GetSdk()
		return ctx, nil
	})

	ctx.Step(`^Given that user-(\d+) initiates a transfer to user-(\d+)$`, func() error {

		sdk.GetStateRootHashLatest(context.Background())

		return utils.Pass
	})

	ctx.Step(`^the transfer amount is"([^"]*)"$`, func() error {
		return utils.Pass
	})

	ctx.Step(`^the transfer gas price is (\+d)$`, func(gasPrice int) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy is given a ttl of (\d+)$`, func(peercount int) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy is put on chain "([^"]*)"$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy response contains a valid deploy hash of length (\d+) and an API version "([^"]*)"$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^wait for a block added event with a timeout of (\d+) seconds$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^that a Transfer has been successfully deployed$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^a deploy is requested via the info_get_deploy RCP method$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^ the deploy data has an API version of "([^"]*)"$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy execution result has "lastBlockAdded" block hash$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy execution has a cost of (\d+) motes$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy has a payment amount of (\d+)$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy has a valid hash$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy has a valid timestamp$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy has a valid body hash$`, func(networkName string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy has a session type of "([^"]*)"$`, func(sessionType string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy is approved by user-(\d+)$`, func(userId int) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy has a gas price of (\d+)$`, func(gasPrice int) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy has a ttl of (\d+)m$`, func(ttl int) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy session has a "([^"]*)" argument value of type "([^"]*)"$`, func(name string, valueTye string) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy session has a "([^"]*)" argument with a numeric value of (\d+)$`, func(name string, value int) error {
		return utils.Pass
	})

	ctx.Step(`^the deploy session has a "([^"]*)" argument with the public key of user-(\d+)$`, func(name string, userId int) error {
		return utils.Pass
	})
}
