package lds

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry/cf-acceptance-tests/helpers"
)

const (
	DEFAULT_TIMEOUT = 30 * time.Second
	CF_PUSH_TIMEOUT = 2 * time.Minute
)

var context helpers.SuiteContext

func TestApplications(t *testing.T) {
	RegisterFailHandler(Fail)

	config := helpers.LoadConfig()
	context = helpers.NewContext(config)
	environment := helpers.NewEnvironment(context)

	BeforeSuite(func() {
		environment.Setup()

		Expect(cf.Cf("push", "lds-dora", "-p", helpers.NewAssets().Dora).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))
		Expect(helpers.CurlAppRoot("lds-dora")).To(ContainSubstring("Hi, I'm Dora!"))
	})

	AfterSuite(func() {
		Expect(cf.Cf("delete", "lds-dora", "-f").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
		environment.Teardown()
	})

	componentName := "LDS"

	rs := []Reporter{}

	if config.ArtifactsDirectory != "" {
		helpers.EnableCFTrace(config, componentName)
		rs = append(rs, helpers.NewJUnitReporter(config, componentName))
	}

	RunSpecsWithDefaultAndCustomReporters(t, componentName, rs)
}
