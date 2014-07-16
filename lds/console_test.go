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

var _ = Describe("Console Tests", func() {
		
	BeforeEach(func() {
		Expect(cf.Cf("push", "lds-java", "-p", helpers.NewAssets().Java).Wait(CF_PUSH_TIMEOUT)).To(Exit(0))
		Expect(helpers.CurlAppRoot("lds-java")).To(ContainSubstring("Hello, from your friendly neighborhood Java JSP!"))
	})

	AfterEach(func() {
		Expect(cf.Cf("delete", "lds-java", "-f").Wait(DEFAULT_TIMEOUT)).To(Exit(0))
	})
		

	//TODO deploy an app that listens on the console port and actually sends data that can be validated.
	It("successfully connect to a console tunnel", func() {
		tunnel := cf.Cf("console-tunnel", "lds-java")
		defer tunnel.Kill()
		time.Sleep(5 * time.Second)
		Expect(tunnel).To(Say("localhost:1830"))

		conn, err := net.DialTimeout("tcp", "localhost:1830", 5*time.Second)
		Expect(err).To(BeNil())
		defer conn.Close()

		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		one := make([]byte, 1)
		_, err = conn.Read(one)
		Expect(err).To(BeNil())
		Expect(tunnel).To(Say("Connected"))
	})
})
