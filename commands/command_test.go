package commands_test

import (
	"github.com/bakku/gom/commands"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Command", func() {
	var _ = Describe("#Select", func() {
		It("should return a generator for generate", func() {
			command, err := commands.Select("generate")

			var c *commands.Generator

			Expect(err).NotTo(HaveOccurred())
			Expect(command).To(BeAssignableToTypeOf(c))
		})

		It("should return a migrator for migrate", func() {
			command, err := commands.Select("migrate")

			var c *commands.Migrator

			Expect(err).NotTo(HaveOccurred())
			Expect(command).To(BeAssignableToTypeOf(c))
		})

		It("should return an initializer for init", func() {
			command, err := commands.Select("init")

			var c *commands.Initializer

			Expect(err).NotTo(HaveOccurred())
			Expect(command).To(BeAssignableToTypeOf(c))
		})

		It("should return a backroller for rollback", func() {
			command, err := commands.Select("rollback")

			var c *commands.Backroller

			Expect(err).NotTo(HaveOccurred())
			Expect(command).To(BeAssignableToTypeOf(c))
		})

		It("should return an error if command does not exist", func() {
			command, err := commands.Select("wrong")

			Expect(command).To(BeNil())
			Expect(err).NotTo(BeNil())
		})
	})
})
