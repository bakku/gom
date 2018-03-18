package commands_test

import (
	"errors"
	"os"

	"github.com/bakku/gom/commands"
	"github.com/bakku/gom/mocks"
	"github.com/bakku/gom/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rollback", func() {
	BeforeEach(func() {
		var err error

		db, err = util.InitDB()
		if err != nil {
			panic(err)
		}

		db.Exec("CREATE TABLE IF NOT EXISTS schema_migrations (migration CHAR(14));")
	})

	AfterEach(func() {
		db.Exec("DROP TABLE IF EXISTS schema_migrations")
		db.Exec("DROP TABLE IF EXISTS posts")
	})

	Describe("Run", func() {
		It("should return no error if no migration was migrated", func() {
			command := commands.Backroller{
				DB: db,
			}

			err := command.Run()

			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an error if migrations dir does not exist", func() {
			db.Exec("INSERT INTO schema_migrations VALUES('20180101100000')")

			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{false}

			command := &commands.Backroller{
				DB:             db,
				FileDirChecker: fileDirChecker,
			}

			err := command.Run()
			Expect(err).To(HaveOccurred())
		})

		It("should return an error if migrations dir could not be read", func() {
			db.Exec("INSERT INTO schema_migrations VALUES('20180101100000')")

			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			dirReader := &mocks.DirReader{}
			dirReader.ReadCall.Returns.Error = errors.New("err")

			command := &commands.Backroller{
				DB:             db,
				FileDirChecker: fileDirChecker,
				DirReader:      dirReader,
			}

			err := command.Run()
			Expect(err).To(HaveOccurred())
		})

		It("should return an error if it cannot find the migration files among the available migrations", func() {
			db.Exec("INSERT INTO schema_migrations VALUES('20180101100000')")

			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			dirReader := &mocks.DirReader{}
			dirReader.ReadCall.Returns.DirSlice = []os.FileInfo{
				&mocks.FileInfo{
					N: "20180607142555_create_posts_table",
				},
				&mocks.FileInfo{
					N: "20180609113232_create_comments_table",
				},
			}

			command := &commands.Backroller{
				DB:             db,
				FileDirChecker: fileDirChecker,
				DirReader:      dirReader,
			}

			err := command.Run()
			Expect(err).To(HaveOccurred())
		})

		It("should return an error if it cannot read the migration file", func() {
			db.Exec("INSERT INTO schema_migrations VALUES('20180101100000')")

			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			dirReader := &mocks.DirReader{}
			dirReader.ReadCall.Returns.DirSlice = []os.FileInfo{
				&mocks.FileInfo{
					N: "20180101100000_create_posts_table",
				},
			}

			fileReader := &mocks.FileReader{}
			fileReader.ReadCall.Returns.String = []string{""}
			fileReader.ReadCall.Returns.Error = []error{errors.New("hi")}

			command := &commands.Backroller{
				DB:             db,
				FileDirChecker: fileDirChecker,
				DirReader:      dirReader,
				FileReader:     fileReader,
			}

			err := command.Run()

			Expect(fileReader.ReadCall.Receives.Path[0]).To(Equal("db/migrations/20180101100000_create_posts_table/down.sql"))
			Expect(err).To(HaveOccurred())
		})

		It("should return no error if ran correctly", func() {
			db.Exec("INSERT INTO schema_migrations VALUES('20180101100000')")
			db.Exec("CREATE TABLE posts(id serial, content text) ;")

			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			dirReader := &mocks.DirReader{}
			dirReader.ReadCall.Returns.DirSlice = []os.FileInfo{
				&mocks.FileInfo{
					N: "20180101100000_create_posts_table",
				},
			}

			fileReader := &mocks.FileReader{}
			fileReader.ReadCall.Returns.String = []string{"DROP TABLE posts;"}
			fileReader.ReadCall.Returns.Error = []error{nil}

			command := &commands.Backroller{
				DB:             db,
				FileDirChecker: fileDirChecker,
				DirReader:      dirReader,
				FileReader:     fileReader,
			}

			err := command.Run()
			Expect(err).NotTo(HaveOccurred())

			var exists bool

			err = db.QueryRow("SELECT EXISTS ( SELECT * FROM schema_migrations WHERE migration = '20180101100000' );").Scan(&exists)
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(Equal(false))

			err = db.QueryRow("SELECT EXISTS ( SELECT 1 FROM information_schema.tables WHERE table_name = 'posts' );").Scan(&exists)
			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(Equal(false))
		})
	})
})
