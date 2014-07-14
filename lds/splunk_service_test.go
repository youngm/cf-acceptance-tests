package lds

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry/cf-acceptance-tests/lds/helpers"
)

var _ = Describe("Splunk Service Tests", func() {

	It("successfully create and bind a Splunk service", func() {
		instanceName := "splunk-service-basic"

		service := cf.Cf("create-splunk-service", instanceName, "ics.esm").Wait(DEFAULT_TIMEOUT)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))
		
		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})
})