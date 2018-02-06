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
			dirCreator := &mocks.DirCreator{}
			fileCreator := &mocks.FileCreator{}

			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = false

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				DirCreator: dirCreator,
				FileCreator: fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).To(HaveOccurred())
		})

		It("should try to create the correct directory", func() {
			dirCreator := &mocks.DirCreator{}

			fileCreator := &mocks.FileCreator{}
			fileCreator.FileCreateCall.Returns.Errors.OnCall = -1
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				DirCreator:  dirCreator,
				FileCreator: fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).NotTo(HaveOccurred())
			Expect(dirCreator.DirCreateCall.Receives.Path).
				To(HaveLen(1))

			Expect(dirCreator.DirCreateCall.Receives.Path[0]).
				To(MatchRegexp("^db/migrations/\\d{14}_create_users_table$"))
		})

		It("should return an error if an error occurred during the folder creation", func() {
			dirCreator := &mocks.DirCreator{}
			dirCreator.DirCreateCall.Returns.Error = errors.New("error")
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				DirCreator: dirCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).To(HaveOccurred())
		})

		It("should try to create up.sql and down.sql", func() {
			dirCreator := &mocks.DirCreator{}

			fileCreator := &mocks.FileCreator{}
			fileCreator.FileCreateCall.Returns.Errors.OnCall = -1
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				DirCreator:  dirCreator,
				FileCreator: fileCreator,
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
			dirCreator := &mocks.DirCreator{}

			fileCreator := &mocks.FileCreator{}
			fileCreator.FileCreateCall.Returns.Errors.OnCall = 1
			fileCreator.FileCreateCall.Returns.Errors.Error = errors.New("error")
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				DirCreator:  dirCreator,
				FileCreator: fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).To(HaveOccurred())
		})

		It("should return an error if an error occured during down.sql creation", func() {
			dirCreator := &mocks.DirCreator{}

			fileCreator := &mocks.FileCreator{}
			fileCreator.FileCreateCall.Returns.Errors.OnCall = 2
			fileCreator.FileCreateCall.Returns.Errors.Error = errors.New("error")
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = true

			command := commands.Generator{
				FileDirChecker: fileDirChecker,
				DirCreator:  dirCreator,
				FileCreator: fileCreator,
			}

			err := command.Run("create_users_table")

			Expect(err).To(HaveOccurred())
		})
	})
})
