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

var _ = Describe("Migrate", func() {
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
		db.Exec("DROP TABLE IF EXISTS comments")
	})

	Describe("#Run", func() {
		It("should return an error if migrations dir does not exist", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{false}

			command := &commands.Migrator{
				FileDirChecker: fileDirChecker,
			}

			err := command.Run()
			Expect(err).To(HaveOccurred())
		})

		It("should return an error if migration directory could not be read", func() {
			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{true}

			dirReader := &mocks.DirReader{}
			dirReader.ReadCall.Returns.Error = errors.New("error")

			command := &commands.Migrator{
				FileDirChecker: fileDirChecker,
				DirReader:      dirReader,
			}

			err := command.Run()
			Expect(err).To(HaveOccurred())
		})

		It("should return an error if migration file does not exist", func() {
			db.Exec("INSERT INTO schema_migrations (migration) VALUES ('20180101000000');")

			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{true, true, false}

			dirReader := &mocks.DirReader{}
			dirReader.ReadCall.Returns.DirSlice = []os.FileInfo{
				&mocks.FileInfo{
					N: "20180101000000_create_users_table",
				},
				&mocks.FileInfo{
					N: "20180607142555_create_posts_table",
				},
				&mocks.FileInfo{
					N: "20180609113232_create_comments_table",
				},
			}

			fileReader := &mocks.FileReader{}
			fileReader.ReadCall.Returns.String = []string{"SELECT * FROM schema_migrations"}
			fileReader.ReadCall.Returns.Error = []error{nil}

			command := &commands.Migrator{
				FileDirChecker: fileDirChecker,
				DirReader:      dirReader,
				DB:             db,
				FileReader:     fileReader,
			}

			err := command.Run()

			Expect(len(fileDirChecker.FileDirExistsCall.Receives.Path)).
				To(Equal(3))

			Expect(fileDirChecker.FileDirExistsCall.Receives.Path[1]).
				To(Equal("db/migrations/20180607142555_create_posts_table/up.sql"))
			Expect(fileDirChecker.FileDirExistsCall.Receives.Path[2]).
				To(Equal("db/migrations/20180609113232_create_comments_table/up.sql"))

			Expect(err).To(HaveOccurred())
		})

		It("should return an error if it could not read migration file", func() {
			db.Exec("INSERT INTO schema_migrations (migration) VALUES ('20180101000000');")

			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{true, true, true}

			dirReader := &mocks.DirReader{}
			dirReader.ReadCall.Returns.DirSlice = []os.FileInfo{
				&mocks.FileInfo{
					N: "20180101000000_create_users_table",
				},
				&mocks.FileInfo{
					N: "20180607142555_create_posts_table",
				},
				&mocks.FileInfo{
					N: "20180609113232_create_comments_table",
				},
			}

			fileReader := &mocks.FileReader{}
			fileReader.ReadCall.Returns.String = []string{
				"BEGIN; CREATE TABLE posts ( id SERIAL, title TEXT ) ; COMMIT;",
				"BEGIN; CREATE TABLE comments ( id SERIAL, comment TEXT ) ; COMMIT;",
			}
			fileReader.ReadCall.Returns.Error = []error{nil, errors.New("error")}

			command := &commands.Migrator{
				FileDirChecker: fileDirChecker,
				DirReader:      dirReader,
				DB:             db,
				FileReader:     fileReader,
			}

			err := command.Run()

			Expect(len(fileReader.ReadCall.Receives.Path)).
				To(Equal(2))

			Expect(fileReader.ReadCall.Receives.Path[0]).
				To(Equal("db/migrations/20180607142555_create_posts_table/up.sql"))
			Expect(fileReader.ReadCall.Receives.Path[1]).
				To(Equal("db/migrations/20180609113232_create_comments_table/up.sql"))

			Expect(err).To(HaveOccurred())
		})

		It("should run the migrations", func() {
			db.Exec("INSERT INTO schema_migrations (migration) VALUES ('20180101000000');")

			fileDirChecker := &mocks.FileDirChecker{}
			fileDirChecker.FileDirExistsCall.Returns.Bool = []bool{true, true, true}

			dirReader := &mocks.DirReader{}
			dirReader.ReadCall.Returns.DirSlice = []os.FileInfo{
				&mocks.FileInfo{
					N: "20180101000000_create_users_table",
				},
				&mocks.FileInfo{
					N: "20180202123456_create_posts_table",
				},
				&mocks.FileInfo{
					N: "20180203123456_create_comments_table",
				},
			}

			fileReader := &mocks.FileReader{}
			fileReader.ReadCall.Returns.String = []string{
				"BEGIN; CREATE TABLE posts ( id SERIAL, title TEXT ) ; COMMIT;",
				"BEGIN; CREATE TABLE comments ( id SERIAL, comment TEXT ) ; COMMIT;",
			}

			fileReader.ReadCall.Returns.Error = []error{nil, nil}

			command := commands.Migrator{
				FileDirChecker: fileDirChecker,
				DirReader:      dirReader,
				DB:             db,
				FileReader:     fileReader,
			}

			err := command.Run()
			Expect(err).NotTo(HaveOccurred())

			var exists bool

			err = db.QueryRow("SELECT EXISTS ( SELECT 1 FROM information_schema.tables WHERE table_name = 'posts' );").Scan(&exists)

			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(Equal(true))

			err = db.QueryRow("SELECT EXISTS ( SELECT 1 FROM information_schema.tables WHERE table_name = 'comments' );").Scan(&exists)

			Expect(err).NotTo(HaveOccurred())
			Expect(exists).To(Equal(true))
		})
	})
})
