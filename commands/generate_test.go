package commands_test

import (
	"errors"

	"github.com/bakku/gom/commands"
	"github.com/bakku/gom/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generator", func() {
	Describe("#Run", func() {
		It("should return an error if no migration name was passed", func() {
			command := commands.Generator{}
			err := command.Run()

			Expect(err).To(HaveOccurred())
		})

		It("should return an error if schema.sql file does not exist", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = false

			fileAppender := &mocks.FileAppender{}

			dirCreator := &mocks.DirCreator{}
			fileCreator := &mocks.FileCreator{}

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				FileAppender: fileAppender,
				DirCreator:     dirCreator,
				FileCreator:    fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).To(HaveOccurred())
		})

		It("should append new migration insert into schema.sql file", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			fileAppender := &mocks.FileAppender{}

			dirCreator := &mocks.DirCreator{}
			fileCreator := &mocks.FileCreator{}


			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				FileAppender:   fileAppender,
				DirCreator:     dirCreator,
				FileCreator:    fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).NotTo(HaveOccurred())
			Expect(fileAppender.AppendCall.Receives.Content).
				To(MatchRegexp("INSERT INTO schema_migrations VALUES \\(\"\\d{14}\"\\) ;\n"))
		})

		It("should return an error if schema append returns an error", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			fileAppender := &mocks.FileAppender{}
			fileAppender.AppendCall.Returns.Error = errors.New("error")

			dirCreator := &mocks.DirCreator{}
			fileCreator := &mocks.FileCreator{}


			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				FileAppender:   fileAppender,
				DirCreator:     dirCreator,
				FileCreator:    fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).To(HaveOccurred())
		})

		It("should try to create the correct directory", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			fileAppender := &mocks.FileAppender{}

			dirCreator := &mocks.DirCreator{}

			fileCreator := &mocks.FileCreator{}
			fileCreator.FileCreateCall.Returns.Errors.OnCall = -1


			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				FileAppender: fileAppender,
				DirCreator:     dirCreator,
				FileCreator:    fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).NotTo(HaveOccurred())
			Expect(dirCreator.DirCreateCall.Receives.Path).
				To(HaveLen(1))

			Expect(dirCreator.DirCreateCall.Receives.Path[0]).
				To(MatchRegexp("^db/migrations/\\d{14}_create_users_table$"))
		})

		It("should return an error if an error occurred during the folder creation", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			fileAppender := &mocks.FileAppender{}

			dirCreator := &mocks.DirCreator{}
			dirCreator.DirCreateCall.Returns.Error = errors.New("error")

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				FileAppender: fileAppender,
				DirCreator:     dirCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).To(HaveOccurred())
		})

		It("should try to create up.sql and down.sql", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			fileAppender := &mocks.FileAppender{}

			dirCreator := &mocks.DirCreator{}

			fileCreator := &mocks.FileCreator{}
			fileCreator.FileCreateCall.Returns.Errors.OnCall = -1


			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				FileAppender: fileAppender,
				DirCreator:     dirCreator,
				FileCreator:    fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).NotTo(HaveOccurred())

			Expect(fileCreator.ReceivedPaths).To(HaveLen(2))
			Expect(fileCreator.ReceivedPaths[0]).
				To(MatchRegexp("^db/migrations/\\d{14}_create_users_table/up.sql$"))

			Expect(fileCreator.ReceivedPaths[1]).
				To(MatchRegexp("^db/migrations/\\d{14}_create_users_table/down.sql$"))
		})

		It("should return an error if an error occured during up.sql creation", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			fileAppender := &mocks.FileAppender{}

			dirCreator := &mocks.DirCreator{}

			fileCreator := &mocks.FileCreator{}
			fileCreator.FileCreateCall.Returns.Errors.OnCall = 1
			fileCreator.FileCreateCall.Returns.Errors.Error = errors.New("error")

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				FileAppender: fileAppender,
				DirCreator:     dirCreator,
				FileCreator:    fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).To(HaveOccurred())
		})

		It("should return an error if an error occured during down.sql creation", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			fileAppender := &mocks.FileAppender{}

			dirCreator := &mocks.DirCreator{}

			fileCreator := &mocks.FileCreator{}
			fileCreator.FileCreateCall.Returns.Errors.OnCall = 2
			fileCreator.FileCreateCall.Returns.Errors.Error = errors.New("error")

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				FileAppender: fileAppender,
				DirCreator:     dirCreator,
				FileCreator:    fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).To(HaveOccurred())
		})
	})
})
