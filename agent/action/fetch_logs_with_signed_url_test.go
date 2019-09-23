package action_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-agent/agent/action"
	fakeblobprovider "github.com/cloudfoundry/bosh-agent/agent/http_blob_provider/http_blob_providerfakes"
	boshdirs "github.com/cloudfoundry/bosh-agent/settings/directories"
	boshassert "github.com/cloudfoundry/bosh-utils/assert"
	boshcrypto "github.com/cloudfoundry/bosh-utils/crypto"
	fakecmd "github.com/cloudfoundry/bosh-utils/fileutil/fakes"
)

var _ = Describe("FetchLogsWithSignedURLAction", func() {
	var (
		compressor           *fakecmd.FakeCompressor
		copier               *fakecmd.FakeCopier
		dirProvider          boshdirs.Provider
		fakeHTTPBlobProvider *fakeblobprovider.FakeHTTPBlobProvider
		action               FetchLogsWithSignedURLAction
	)

	BeforeEach(func() {
		compressor = fakecmd.NewFakeCompressor()
		dirProvider = boshdirs.NewProvider("/fake/dir")
		copier = fakecmd.NewFakeCopier()
		fakeHTTPBlobProvider = &fakeblobprovider.FakeHTTPBlobProvider{}
		action = NewFetchLogsWithSignedURLAction(compressor, copier, dirProvider, fakeHTTPBlobProvider)
	})

	AssertActionIsAsynchronous(action)
	AssertActionIsNotPersistent(action)
	AssertActionIsLoggable(action)

	AssertActionIsNotResumable(action)
	AssertActionIsNotCancelable(action)

	Describe("Run", func() {
		testLogs := func(logType string, filters []string, expectedFilters []string) {
			copier.FilteredCopyToTempTempDir = "/fake-temp-dir"
			compressor.CompressFilesInDirTarballPath = "logs_test.tar"

			multidigestSha := boshcrypto.MustNewMultipleDigest(boshcrypto.NewDigest(boshcrypto.DigestAlgorithmSHA1, "sec_dep_sha1"))
			sha1 := multidigestSha.String()
			fakeHTTPBlobProvider.UploadReturns(multidigestSha, nil)

			logs, err := action.Run(FetchLogsWithSignedURLRequest{SignedURL: "foobar", LogType: logType, Filters: filters})
			Expect(err).ToNot(HaveOccurred())

			var expectedPath string
			switch logType {
			case "job":
				expectedPath = filepath.Join("/fake", "dir", "sys", "log")
			case "agent":
				expectedPath = filepath.Join("/fake", "dir", "bosh", "log")
			}

			Expect(copier.FilteredCopyToTempDir).To(boshassert.MatchPath(expectedPath))
			Expect(copier.FilteredCopyToTempFilters).To(Equal(expectedFilters))

			Expect(copier.FilteredCopyToTempTempDir).To(Equal(compressor.CompressFilesInDirDir))
			Expect(copier.CleanUpTempDir).To(Equal(compressor.CompressFilesInDirDir))

			actualSignedURL, actualTarballPath := fakeHTTPBlobProvider.UploadArgsForCall(0)
			Expect(actualSignedURL).To(Equal("foobar"))
			Expect(actualTarballPath).To(Equal(compressor.CompressFilesInDirTarballPath))

			boshassert.MatchesJSONString(GinkgoT(), logs, `{"sha1":"`+sha1+`"}`)
		}

		It("logs errs if given invalid log type", func() {
			_, err := action.Run(FetchLogsWithSignedURLRequest{LogType: "other-logs", Filters: []string{}})
			Expect(err).To(HaveOccurred())
		})

		It("agent logs with filters", func() {
			filters := []string{"**/*.stdout.log", "**/*.stderr.log"}
			expectedFilters := []string{"**/*.stdout.log", "**/*.stderr.log"}
			testLogs("agent", filters, expectedFilters)
		})

		It("agent logs without filters", func() {
			filters := []string{}
			expectedFilters := []string{"**/*"}
			testLogs("agent", filters, expectedFilters)
		})

		It("job logs without filters", func() {
			filters := []string{}
			expectedFilters := []string{"**/*"}
			testLogs("job", filters, expectedFilters)
		})

		It("job logs with filters", func() {
			filters := []string{"**/*.stdout.log", "**/*.stderr.log"}
			expectedFilters := []string{"**/*.stdout.log", "**/*.stderr.log"}
			testLogs("job", filters, expectedFilters)
		})

		It("cleans up compressed package after uploading it to blobstore", func() {
			var beforeCleanUpTarballPath, afterCleanUpTarballPath string

			compressor.CompressFilesInDirTarballPath = "/fake-compressed-logs.tar"

			_, err := action.Run(FetchLogsWithSignedURLRequest{LogType: "job", Filters: []string{}})
			Expect(err).ToNot(HaveOccurred())

			// Logs are not cleaned up before blobstore upload
			Expect(beforeCleanUpTarballPath).To(Equal(""))

			// Deleted after it was uploaded
			afterCleanUpTarballPath = compressor.CleanUpTarballPath
			Expect(afterCleanUpTarballPath).To(Equal("/fake-compressed-logs.tar"))
		})
	})
})
