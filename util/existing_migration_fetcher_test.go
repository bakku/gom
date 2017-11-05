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
	Describe("fetch", func() {
		Context("If no error occurs", func() {
			It("should return the existing migrations", func() {
				mock := &mocks.MockDirReader{}
				mock.ReadCall.Returns.DirSlice = []os.FileInfo{
					&mocks.MockFileInfo{
						N: "20170405135505_create_users_table",
					},
					&mocks.MockFileInfo{
						N: "20170607142555_create_posts_table",
					},
					&mocks.MockFileInfo{
						N: "20170609113232_create_comments_table",
					},
				}
				mock.ReadCall.Returns.Error = nil

				fetcher := &util.ExistingMigrationFetcher{
					DirReader: mock,
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
				mock := &mocks.MockDirReader{}
				mock.ReadCall.Returns.DirSlice = []os.FileInfo{}
				mock.ReadCall.Returns.Error = errors.New("error")

				fetcher := &util.ExistingMigrationFetcher{
					DirReader: mock,
				}

				actual, err := fetcher.Fetch()
				expected := []string{}

				Expect(err).NotTo(BeNil())
				Expect(actual).To(Equal(expected))
			})
		})
	})
})
