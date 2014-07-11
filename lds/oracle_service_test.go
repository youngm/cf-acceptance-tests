package lds

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry/cf-acceptance-tests/lds/helpers"
)

var _ = Describe("Oracle Service Tests", func() {

	It("successfully create and bind an existing service", func() {
		instanceName := "oracle-service-existing"

		service := cf.Cf("create-oracle-service", instanceName, "Existing", "-s", "a047", "-u", "COLLECTOR_DEV1", "-p", "f8XkKgD3R7ywDld").Wait(DEFAULT_TIMEOUT)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))
		
		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})

	It("successfully create and bind a test service", func() {
		instanceName := "oracle-service-test"

		service := cf.Cf("create-oracle-service", instanceName, "Test Oracle 11g").Wait(DEFAULT_TIMEOUT)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))

		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})
	
	It("successfully create and bind a test service", func() {
		instanceName := "oracle-service-prod"

		service := cf.Cf("create-oracle-service", instanceName, "Prod Oracle 11g").Wait(DEFAULT_TIMEOUT)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))

		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})
})