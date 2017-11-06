package commands_test

import (
	"github.com/bakku/gom/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rollback", func() {
	Describe("Run", func() {
		It("should return no error if ran correctly", func() {
			command := commands.Backroller{}

			err := command.Run()

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
