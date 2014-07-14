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