package lds

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry/cf-acceptance-tests/lds/helpers"
)

var _ = Describe("LDS Account Service Tests", func() {

	It("successfully create and bind a simple lds account service", func() {
		instanceName := "lds-account-service-stage"

		service := cf.Cf("create-lds-account-service", instanceName, "stage", "LDAP-PaaS", "-p", "292013TueAD31").Wait(DEFAULT_TIMEOUT)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))
		
		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})
})