package diego

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/generator"
	"github.com/cloudfoundry/cf-acceptance-tests/helpers"
)

var _ = Describe("An application being staged with Diego", func() {
	var appName string

	BeforeEach(func() {
		appName = generator.RandomName()
	})

	AfterEach(func() {
		Eventually(cf.Cf("delete", appName, "-f"), DEFAULT_TIMEOUT).Should(Exit(0))
	})

	It("has its staging log streamed during a push", func() {
		//Diego needs a custom buildpack until the ruby buildpack lands
		Eventually(cf.Cf("push", appName, "-p", helpers.NewAssets().Dora, "--no-start", "-b=https://github.com/cloudfoundry/cf-buildpack-ruby/archive/master.zip"), CF_PUSH_TIMEOUT).Should(Exit(0))
		Eventually(cf.Cf("set-env", appName, "CF_DIEGO_BETA", "true"), DEFAULT_TIMEOUT).Should(Exit(0))

		start := cf.Cf("start", appName)

		Eventually(start, CF_PUSH_TIMEOUT).Should(Say("Downloading App Package"))
		Eventually(start, CF_PUSH_TIMEOUT).Should(Say("Downloaded App Package"))
		Eventually(start, CF_PUSH_TIMEOUT).Should(Say(`Staging\.\.\.`))
		Eventually(start, CF_PUSH_TIMEOUT).Should(Exit(0))
	})
})
