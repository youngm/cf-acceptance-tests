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

var _ = Describe("App Dynamics Service Tests", func() {
	smInstanceName := "servivcemanager-service-for-ad"

	BeforeEach(func() {
		service := cf.Cf("create-servicemanager-service", smInstanceName, "test-cf", "dev", "true").Wait(300 * time.Second)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(smInstanceName))
		
		bind := cf.Cf("bind-service", "lds-dora", smInstanceName).Wait(300 * time.Second)
		Expect(bind).To(Exit(0))
		Expect(bind).To(Say(smInstanceName))
	})

	AfterEach(func() {
		unbind := cf.Cf("unbind-service", "lds-dora", smInstanceName).Wait(DEFAULT_TIMEOUT)
		Expect(unbind).To(Exit(0))
		Expect(unbind).To(Say(smInstanceName))
	
		service := cf.Cf("delete-service", smInstanceName, "-f").Wait(DEFAULT_TIMEOUT)
		Expect(service).To(Exit(0))

		services := cf.Cf("services").Wait(DEFAULT_TIMEOUT)
		Expect(services).To(Exit(0))
		Expect(services.Out.Contents()).NotTo(ContainSubstring(smInstanceName))
	})

	It("successfully create and bind an app dynamics service", func() {
		instanceName := "app-dynamics-service-saas-non-prod"

		service := cf.Cf("create-app-dynamics-service", instanceName, "saas-non-prod").Wait(DEFAULT_TIMEOUT)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))

		helpers.TestAndCleanupService(instanceName, DEFAULT_TIMEOUT)
	})

})