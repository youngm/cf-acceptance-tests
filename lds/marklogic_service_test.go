package lds

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry/cf-acceptance-tests/lds/helpers"
)

var _ = Describe("Marklogic Service Tests", func() {

	It("successfully create and bind an existing service", func() {
		instanceName := "marklogic-service-existing"

		service := cf.Cf("create-marklogic-service", instanceName, "Existing", "-a", "labslcl0622.qalab.ldschurch.org", "-n", "9007", "-u", "ml-deleteme_19e641f0-1440-4ec0-8e34-dfb839c0dc1c_2014-05-01-19-33-06-user", "-p", "ml-deleteme_19e641f0-1440-4ec0-8e34-dfb839c0dc1c_2014-05-01-19-33-063024pass").Wait(DEFAULT_TIMEOUT)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))
		
		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})

	It("successfully create and bind a test service", func() {
		instanceName := "marklogic-service-test"

		service := cf.Cf("create-marklogic-service", instanceName, "Test Provisioning").Wait(60 * time.Second)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))

		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})
	
	It("successfully create and bind a prod service", func() {
		instanceName := "marklogic-service-prod"

		service := cf.Cf("create-marklogic-service", instanceName, "Default").Wait(60 * time.Second)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))

		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})
})