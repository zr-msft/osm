package tresor

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/openservicemesh/osm/pkg/certificate"
)

var _ = Describe("Test Tresor Debugger", func() {
	Context("test ListIssuedCertificates()", func() {
		// Setup:
		//   1. Create a new (fake) certificate
		//   2. Reuse the same certificate as the Issuing CA
		//   3. Populate the CertManager's cache w/ cert
		cert := NewFakeCertificate()
		cm := CertManager{}
		cm.cache.Store("foo", cert)

		It("lists all issued certificates", func() {
			actual := cm.ListIssuedCertificates()
			expected := []*certificate.Certificate{cert}
			Expect(actual).To(Equal(expected))
		})
	})
})
