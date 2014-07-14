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

var _ = Describe("Service Manager Service Tests", func() {

	It("successfully create and bind a servicemanager service", func() {
		instanceName := "servivcemanager-service-new"
		timeTag := time.Now().Format("2006_01_02-15h04m05.999s")
		
		service := cf.Cf("create-servicemanager-service", instanceName, "CF_TEST_"+timeTag, "dev", "false", "-d", "1 801 240 5555", "-c", "NIELSENSK", "-m", "NIELSENSK", "-g", "CHQ-ENG-PLATFORMS MIDDLEWARE", "-p", "5013401", "-o", "1-Low").Wait(300 * time.Second)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))
		
		helpers.TestAndCleanupService(instanceName, 300 * time.Second)
	})

	It("successfully create and bind an existing servicemanager service", func() {
		instanceName := "servivcemanager-service-existing"

		service := cf.Cf("create-servicemanager-service", instanceName, "test-cf", "dev", "true").Wait(300 * time.Second)
		Expect(service).To(Exit(0))
		Expect(service).To(Say(instanceName))

		helpers.TestAndCleanupService(instanceName, 300 * time.Second)
	})

})