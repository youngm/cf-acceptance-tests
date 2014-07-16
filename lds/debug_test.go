package lds

import (
	"net"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry/cf-acceptance-tests/helpers"
)

var _ = Describe("Debug Tests", func() {

	BeforeEach(func() {
		Expect(cf.Cf("push", "lds-java", "-p", helpers.NewAssets().Java).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))
		Expect(helpers.CurlAppRoot("lds-java")).To(ContainSubstring("Hello, from your friendly neighborhood Java JSP!"))
	})

	AfterEach(func() {
		Expect(cf.Cf("delete", "lds-java", "-f").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
	})

	//TODO deploy an app that listens on the debug port and actually sends data that can be validated.
	It("successfully connect to a debug tunnel", func() {
		debug := cf.Cf("debug", "lds-java", "-d", "run").Wait(DEFAULT_TIMEOUT)
		Expect(debug).To(Exit(0))
		Expect(debug).To(Say("lds-java"))
		
		tunnel := cf.Cf("debug-tunnel", "lds-java")
		defer tunnel.Kill()
		time.Sleep(5 * time.Second)
		Expect(tunnel).To(Say("localhost:1847"))

		conn, err := net.DialTimeout("tcp", "localhost:1847", 5*time.Second)
		Expect(err).To(BeNil())
		defer conn.Close()
		time.Sleep(2 * time.Second)
		Expect(tunnel).To(Say("Connected"))
	})
})
