package helpers

import (
	"time"
	
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	
	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
)

func TestAndCleanupService(instanceName string, timeout time.Duration) {
	services := cf.Cf("services").Wait(timeout)
	Expect(services).To(Exit(0))
	Expect(services).To(Say(instanceName))

	bind := cf.Cf("bind-service", "lds-dora", instanceName).Wait(timeout)
	Expect(bind).To(Exit(0))
	Expect(bind).To(Say(instanceName))

	unbind := cf.Cf("unbind-service", "lds-dora", instanceName).Wait(timeout)
	Expect(unbind).To(Exit(0))
	Expect(unbind).To(Say(instanceName))

	service := cf.Cf("delete-service", instanceName, "-f").Wait(timeout)
	Expect(service).To(Exit(0))
	
	services = cf.Cf("services").Wait(timeout)
	Expect(services).To(Exit(0))
	Expect(services.Out.Contents()).NotTo(ContainSubstring(instanceName))
}
