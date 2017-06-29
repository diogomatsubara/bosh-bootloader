package acceptance_test

import (
	acceptance "github.com/cloudfoundry/bosh-bootloader/acceptance-tests"
	"github.com/cloudfoundry/bosh-bootloader/acceptance-tests/actors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("idempotent test", func() {
	var (
		bbl           actors.BBL
		configuration acceptance.Config
	)

	BeforeEach(func() {
		var err error
		configuration, err = acceptance.LoadConfig()
		Expect(err).NotTo(HaveOccurred())

		bbl = actors.NewBBL(configuration.StateFileDir, pathToBBL, configuration, "reentrant-env")
	})

	AfterEach(func() {
		if !CurrentGinkgoTestDescription().Failed {
			bbl.Destroy()
		}
	})

	It("is able to bbl up idempotently with a director", func() {
		bbl.Up(actors.GetIAAS(configuration), []string{"--name", bbl.PredefinedEnvID()})
		bbl.Up(actors.GetIAAS(configuration), []string{})
	})

	It("is able to bbl up idempotently with no director", func() {
		bbl.Up(actors.GetIAAS(configuration), []string{"--name", bbl.PredefinedEnvID(), "--no-director"})
		bbl.Up(actors.GetIAAS(configuration), []string{})
	})
})
