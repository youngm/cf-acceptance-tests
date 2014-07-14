package lds

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry/cf-acceptance-tests/lds/helpers"
)

var _ = Describe("Enterprise WS Service Tests", func() {

	It("successfully create and bind an enterprise ws service", func() {
		instanceName := "enterprise-ws-service-existing"
		service := cf.Cf("create-enterprise-ws-service", instanceName, "int", "cloud-test", "-p", "onlyAccessToServiceNotToMethods").Wait(DEFAULT_TIMEOUT)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))
		
		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})
})