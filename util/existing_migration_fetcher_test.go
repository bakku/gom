package util_test

import (
	"errors"
	"os"

	"github.com/bakku/gom/mocks"
	"github.com/bakku/gom/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExistingMigrationFetcher", func() {
	Describe("Fetch", func() {
		Context("If dir does not exist", func() {
			It("should return an error", func() {
				mockDirReader := &mocks.DirReader{}
				mockDirChecker := &mocks.DirChecker{}
				mockDirChecker.DirExistsCall.Returns.Bool = false

				fetcher := &util.ExistingMigrationFetcher{
					DirReader:  mockDirReader,
					DirChecker: mockDirChecker,
				}

				actual, err := fetcher.Fetch()
				expected := []string{}

				Expect(err).NotTo(BeNil())
				Expect(actual).To(Equal(expected))
			})
		})

		Context("If dir exists", func() {
			Context("If no error occurs", func() {
				It("should return the existing migrations", func() {
					mockDirReader := &mocks.DirReader{}
					mockDirReader.ReadCall.Returns.DirSlice = []os.FileInfo{
						&mocks.FileInfo{
							N: "20170405135505_create_users_table",
						},
						&mocks.FileInfo{
							N: "20170607142555_create_posts_table",
						},
						&mocks.FileInfo{
							N: "20170609113232_create_comments_table",
						},
					}
					mockDirReader.ReadCall.Returns.Error = nil

					mockDirChecker := &mocks.DirChecker{}
					mockDirChecker.DirExistsCall.Returns.Bool = true

					fetcher := &util.ExistingMigrationFetcher{
						DirReader:  mockDirReader,
						DirChecker: mockDirChecker,
					}

					actual, err := fetcher.Fetch()
					expected := []string{
						"20170405135505_create_users_table",
						"20170607142555_create_posts_table",
						"20170609113232_create_comments_table",
					}

					Expect(err).To(BeNil())
					Expect(actual).To(Equal(expected))
				})
			})

			Context("If an error occurs", func() {
				It("should return no migrations and an error", func() {
					mockDirReader := &mocks.DirReader{}
					mockDirReader.ReadCall.Returns.DirSlice = []os.FileInfo{}
					mockDirReader.ReadCall.Returns.Error = errors.New("error")

					mockDirChecker := &mocks.DirChecker{}
					mockDirChecker.DirExistsCall.Returns.Bool = true

					fetcher := &util.ExistingMigrationFetcher{
						DirReader:  mockDirReader,
						DirChecker: mockDirChecker,
					}

					actual, err := fetcher.Fetch()
					expected := []string{}

					Expect(err).NotTo(BeNil())
					Expect(actual).To(Equal(expected))
				})
			})
		})
	})
})
