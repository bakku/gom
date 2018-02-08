package commands_test

import (
	"database/sql"
	"errors"

	"github.com/bakku/gom/commands"
	"github.com/bakku/gom/mocks"
	"github.com/bakku/gom/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/lib/pq"
)

var _ = Describe("Initializer", func() {
	BeforeEach(func() {
		var err error

		db, err = util.InitDB()
		if err != nil {
			panic(err)
		}
	})

	AfterEach(func() {
		db.Exec("DROP TABLE IF EXISTS schema_migrations")
	})

	Describe("#Run", func() {
		It("should create a table called schema_migrations", func() {
			dirCreator := &mocks.DirCreator{}

			fileChecker := &mocks.FileDirChecker{}
			fileChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			fileWriter := &mocks.FileWriter{}

			command := commands.Initializer{
				DB:             db,
				DirCreator:     dirCreator,
				FileDirChecker: fileChecker,
				FileWriter:     fileWriter,
			}

			command.Run()

			var result sql.NullString

			row := db.QueryRow("SELECT to_regclass('schema_migrations')")
			err := row.Scan(&result)

			Expect(err).NotTo(HaveOccurred())
			Expect(result.Valid).To(BeTrue())
		})

		It("should try to create the correct directory", func() {
			dirCreator := &mocks.DirCreator{}

			fileChecker := &mocks.FileDirChecker{}
			fileChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			fileWriter := &mocks.FileWriter{}

			command := commands.Initializer{
				DB:             db,
				DirCreator:     dirCreator,
				FileDirChecker: fileChecker,
				FileWriter:     fileWriter,
			}

			err := command.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(dirCreator.DirCreateCall.Receives.Path).
				To(HaveLen(2))

			Expect(dirCreator.DirCreateCall.Receives.Path[0]).
				To(Equal("db"))
			Expect(dirCreator.DirCreateCall.Receives.Path[1]).
				To(Equal("db/migrations"))
		})

		It("should return an error if an error occurred during the folder creation", func() {
			dirCreator := &mocks.DirCreator{}
			dirCreator.DirCreateCall.Returns.Error = errors.New("error")

			fileChecker := &mocks.FileDirChecker{}
			fileChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			command := commands.Initializer{
				DB:             db,
				DirCreator:     dirCreator,
				FileDirChecker: fileChecker,
			}

			err := command.Run()

			Expect(err).To(HaveOccurred())
		})

		It("should create schema.sql with the base schema if it does not exist yet", func() {
			dirCreator := &mocks.DirCreator{}

			fileChecker := &mocks.FileDirChecker{}
			fileChecker.FileDirExistsCall.Returns.Bool = []bool{false}

			fileWriter := &mocks.FileWriter{}

			command := commands.Initializer{
				DB:             db,
				DirCreator:     dirCreator,
				FileDirChecker: fileChecker,
				FileWriter:     fileWriter,
			}
			err := command.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(fileWriter.WriteCall.Receives.Path).To(Equal("db/schema.sql"))
			Expect(fileWriter.WriteCall.Receives.Content).NotTo(Equal(""))
		})

		It("should not create schema.sql if it exists", func() {
			dirCreator := &mocks.DirCreator{}

			fileChecker := &mocks.FileDirChecker{}
			fileChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			fileWriter := &mocks.FileWriter{}

			command := commands.Initializer{
				DB:             db,
				DirCreator:     dirCreator,
				FileDirChecker: fileChecker,
				FileWriter:     fileWriter,
			}
			err := command.Run()

			Expect(err).NotTo(HaveOccurred())
			Expect(fileWriter.WriteCall.Receives.Path).To(Equal(""))
			Expect(fileWriter.WriteCall.Receives.Content).To(Equal(""))
		})

		It("should return an error if file creation fails", func() {
			dirCreator := &mocks.DirCreator{}

			fileChecker := &mocks.FileDirChecker{}
			fileChecker.FileDirExistsCall.Returns.Bool = []bool{false}

			fileWriter := &mocks.FileWriter{}
			fileWriter.WriteCall.Returns.Error = errors.New("error")

			command := commands.Initializer{
				DB:             db,
				DirCreator:     dirCreator,
				FileDirChecker: fileChecker,
				FileWriter:     fileWriter,
			}
			err := command.Run()

			Expect(err).To(HaveOccurred())
		})

		It("should return no error when ran correctly", func() {
			dirCreator := &mocks.DirCreator{}

			fileChecker := &mocks.FileDirChecker{}
			fileChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			fileWriter := &mocks.FileWriter{}

			command := commands.Initializer{
				DB:             db,
				DirCreator:     dirCreator,
				FileDirChecker: fileChecker,
				FileWriter:     fileWriter,
			}
			err := command.Run()

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
