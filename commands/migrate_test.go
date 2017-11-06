package commands_test

import (
	"github.com/bakku/gom/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Migrate", func() {
	Describe("#Run", func() {
		It("should return no error when ran correctly", func() {
			command := commands.Migrator{}
			err := command.Run()

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
