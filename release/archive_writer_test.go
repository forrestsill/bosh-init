package release_test

import (
	"errors"

	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	fakecmd "github.com/cloudfoundry/bosh-utils/fileutil/fakes"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-init/release"
	fakerel "github.com/cloudfoundry/bosh-init/release/fakes"
	boshjob "github.com/cloudfoundry/bosh-init/release/job"
	boshlic "github.com/cloudfoundry/bosh-init/release/license"
	boshman "github.com/cloudfoundry/bosh-init/release/manifest"
	boshpkg "github.com/cloudfoundry/bosh-init/release/pkg"
	. "github.com/cloudfoundry/bosh-init/release/resource"
)

var _ = Describe("ArchiveWriter", func() {
	var (
		compressor *fakecmd.FakeCompressor
		fs         *fakesys.FakeFileSystem
		writer     ArchiveWriter

		release      *fakerel.FakeRelease
		pkgFpsToSkip []string
	)

	BeforeEach(func() {
		compressor = fakecmd.NewFakeCompressor()
		fs = fakesys.NewFakeFileSystem()
		fs.TempDirDir = "/staging-release"
		logger := boshlog.NewLogger(boshlog.LevelNone)
		writer = NewArchiveWriter(compressor, fs, logger)

		release = &fakerel.FakeRelease{}
		pkgFpsToSkip = nil
	})

	Describe("Write", func() {
		act := func() (string, error) { return writer.Write(release, pkgFpsToSkip) }

		BeforeEach(func() {
			compressor.CompressFilesInDirTarballPath = "/release-archive"
		})

		It("writes out release.MF", func() {
			compressed := false

			release.ManifestReturns(boshman.Manifest{
				Name:               "name",
				Version:            "ver",
				CommitHash:         "commit",
				UncommittedChanges: true,
				Jobs: []boshman.JobRef{
					{
						Name:        "job",
						Version:     "job-version",
						Fingerprint: "job-fp",
						SHA1:        "job-sha1",
					},
				},
				Packages: []boshman.PackageRef{
					{
						Name:         "pkg",
						Version:      "pkg-version",
						Fingerprint:  "pkg-fp",
						SHA1:         "pkg-sha1",
						Dependencies: []string{"pkg1"},
					},
				},
				CompiledPkgs: []boshman.CompiledPackageRef{
					{
						Name:          "cp",
						Version:       "cp-version",
						Fingerprint:   "cp-fp",
						SHA1:          "cp-sha1",
						OSVersionSlug: "cp-os-slug",
						Dependencies:  []string{"pkg1", "pkg2"},
					},
				},
				License: &boshman.LicenseRef{
					Version:     "lic-version",
					Fingerprint: "lic-fp",
					SHA1:        "lic-sha1",
				},
			})

			compressor.CompressFilesInDirCallBack = func() {
				Expect(fs.ReadFileString("/staging-release/release.MF")).To(Equal(`name: name
version: ver
commit_hash: commit
uncommitted_changes: true
jobs:
- name: job
  version: job-version
  fingerprint: job-fp
  sha1: job-sha1
packages:
- name: pkg
  version: pkg-version
  fingerprint: pkg-fp
  sha1: pkg-sha1
  dependencies:
  - pkg1
compiled_packages:
- name: cp
  version: cp-version
  fingerprint: cp-fp
  sha1: cp-sha1
  stemcell: cp-os-slug
  dependencies:
  - pkg1
  - pkg2
license:
  version: lic-version
  fingerprint: lic-fp
  sha1: lic-sha1
`))
				compressed = true
			}

			path, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(path).To(Equal("/release-archive"))

			Expect(compressed).To(BeTrue())
			Expect(fs.FileExists("/staging-release")).To(BeFalse())
		})

		It("returns error if writing release.MF fails", func() {
			fs.WriteFileErrors["/staging-release/release.MF"] = errors.New("fake-err")

			_, err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))
		})

		It("does not include empty 'jobs', 'packages' or 'compiled_packages' directories", func() {
			compressed := false

			compressor.CompressFilesInDirCallBack = func() {
				Expect(fs.FileExists("/staging-release/jobs")).To(BeFalse())
				Expect(fs.FileExists("/staging-release/packages")).To(BeFalse())
				Expect(fs.FileExists("/staging-release/compiled_packages")).To(BeFalse())
				compressed = true
			}

			path, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(path).To(Equal("/release-archive"))

			Expect(compressed).To(BeTrue())
		})

		It("writes out jobs", func() {
			compressed := false

			fs.WriteFileString("/tmp/job1.tgz", "job1-content")
			fs.WriteFileString("/tmp/job2.tgz", "job2-content")

			release.JobsReturns([]*boshjob.Job{
				boshjob.NewJob(NewResourceWithBuiltArchive("job1", "", "/tmp/job1.tgz", "")),
				boshjob.NewJob(NewResourceWithBuiltArchive("job2", "", "/tmp/job2.tgz", "")),
			})

			compressor.CompressFilesInDirCallBack = func() {
				Expect(fs.FileExists("/staging-release/jobs")).To(BeTrue())
				Expect(fs.ReadFileString("/staging-release/jobs/job1.tgz")).To(Equal("job1-content"))
				Expect(fs.ReadFileString("/staging-release/jobs/job2.tgz")).To(Equal("job2-content"))
				compressed = true
			}

			path, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(path).To(Equal("/release-archive"))

			Expect(compressed).To(BeTrue())
		})

		It("returns error if copying job fails", func() {
			fs.CopyFileError = errors.New("fake-err")

			release.JobsReturns([]*boshjob.Job{
				boshjob.NewJob(NewResourceWithBuiltArchive("job1", "", "/tmp/job1.tgz", "")),
			})

			_, err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))

			Expect(fs.FileExists("/staging-release")).To(BeFalse())
		})

		It("writes out all packages when filtering is off (nil)", func() {
			compressed := false

			fs.WriteFileString("/tmp/pkg1.tgz", "pkg1-content")
			fs.WriteFileString("/tmp/pkg2.tgz", "pkg2-content")

			release.PackagesReturns([]*boshpkg.Package{
				boshpkg.NewPackage(NewResourceWithBuiltArchive("pkg1", "", "/tmp/pkg1.tgz", ""), nil),
				boshpkg.NewPackage(NewResourceWithBuiltArchive("pkg2", "", "/tmp/pkg2.tgz", ""), nil),
			})

			compressor.CompressFilesInDirCallBack = func() {
				Expect(fs.FileExists("/staging-release/packages")).To(BeTrue())
				Expect(fs.ReadFileString("/staging-release/packages/pkg1.tgz")).To(Equal("pkg1-content"))
				Expect(fs.ReadFileString("/staging-release/packages/pkg2.tgz")).To(Equal("pkg2-content"))
				compressed = true
			}

			path, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(path).To(Equal("/release-archive"))

			Expect(compressed).To(BeTrue())
		})

		It("excludes filtered out packages when filtering is on", func() {
			compressed := false
			pkgFpsToSkip = append(pkgFpsToSkip, "pkg1-fp")

			fs.WriteFileString("/tmp/pkg1.tgz", "pkg1-content")
			fs.WriteFileString("/tmp/pkg2.tgz", "pkg2-content")

			release.PackagesReturns([]*boshpkg.Package{
				boshpkg.NewPackage(NewResourceWithBuiltArchive("pkg1", "pkg1-fp", "/tmp/pkg1.tgz", ""), nil),
				boshpkg.NewPackage(NewResourceWithBuiltArchive("pkg2", "pkg2-fp", "/tmp/pkg2.tgz", ""), nil),
			})

			compressor.CompressFilesInDirCallBack = func() {
				Expect(fs.FileExists("/staging-release/packages")).To(BeTrue())
				Expect(fs.FileExists("/staging-release/packages/pkg1.tgz")).To(BeFalse())
				Expect(fs.ReadFileString("/staging-release/packages/pkg2.tgz")).To(Equal("pkg2-content"))
				compressed = true
			}

			path, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(path).To(Equal("/release-archive"))

			Expect(compressed).To(BeTrue())
		})

		It("returns error if copying package fails", func() {
			fs.CopyFileError = errors.New("fake-err")

			release.PackagesReturns([]*boshpkg.Package{
				boshpkg.NewPackage(NewResourceWithBuiltArchive("pkg1", "", "/tmp/pkg1.tgz", ""), nil),
			})

			_, err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))

			Expect(fs.FileExists("/staging-release")).To(BeFalse())
		})

		It("writes out all compiled packages when filtering is off (nil)", func() {
			compressed := false

			fs.WriteFileString("/tmp/cp1.tgz", "cp1-content")
			fs.WriteFileString("/tmp/cp2.tgz", "cp2-content")

			release.CompiledPackagesReturns([]*boshpkg.CompiledPackage{
				boshpkg.NewCompiledPackageWithArchive("cp1", "", "", "/tmp/cp1.tgz", "", nil),
				boshpkg.NewCompiledPackageWithArchive("cp2", "", "", "/tmp/cp2.tgz", "", nil),
			})

			compressor.CompressFilesInDirCallBack = func() {
				Expect(fs.FileExists("/staging-release/compiled_packages")).To(BeTrue())
				Expect(fs.ReadFileString("/staging-release/compiled_packages/cp1.tgz")).To(Equal("cp1-content"))
				Expect(fs.ReadFileString("/staging-release/compiled_packages/cp2.tgz")).To(Equal("cp2-content"))
				compressed = true
			}

			path, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(path).To(Equal("/release-archive"))

			Expect(compressed).To(BeTrue())
		})

		It("excludes filtered out compiled packages when filtering is on", func() {
			compressed := false
			pkgFpsToSkip = append(pkgFpsToSkip, "cp1-fp")

			fs.WriteFileString("/tmp/cp1.tgz", "cp1-content")
			fs.WriteFileString("/tmp/cp2.tgz", "cp2-content")

			release.CompiledPackagesReturns([]*boshpkg.CompiledPackage{
				boshpkg.NewCompiledPackageWithArchive("cp1", "cp1-fp", "", "/tmp/cp1.tgz", "", nil),
				boshpkg.NewCompiledPackageWithArchive("cp2", "cp2-fp", "", "/tmp/cp2.tgz", "", nil),
			})

			compressor.CompressFilesInDirCallBack = func() {
				Expect(fs.FileExists("/staging-release/compiled_packages")).To(BeTrue())
				Expect(fs.FileExists("/staging-release/compiled_packages/cp1.tgz")).To(BeFalse())
				Expect(fs.ReadFileString("/staging-release/compiled_packages/cp2.tgz")).To(Equal("cp2-content"))
				compressed = true
			}

			path, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(path).To(Equal("/release-archive"))

			Expect(compressed).To(BeTrue())
		})

		It("returns error if copying compiled package fails", func() {
			fs.CopyFileError = errors.New("fake-err")

			release.CompiledPackagesReturns([]*boshpkg.CompiledPackage{
				boshpkg.NewCompiledPackageWithArchive("cp1", "", "", "/tmp/cp1.tgz", "", nil),
			})

			_, err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))

			Expect(fs.FileExists("/staging-release")).To(BeFalse())
		})

		It("writes out license and unpacks license files into root", func() {
			compressed := false
			decompressed := false

			fs.WriteFileString("/tmp/lic.tgz", "license-content")

			release.LicenseReturns(boshlic.NewLicense(
				NewResourceWithBuiltArchive("lic", "", "/tmp/lic.tgz", "")))

			compressor.DecompressFileToDirCallBack = func() {
				Expect(compressor.DecompressFileToDirTarballPaths).To(Equal([]string{"/tmp/lic.tgz"}))
				Expect(compressor.DecompressFileToDirDirs).To(Equal([]string{"/staging-release"}))
				Expect(compressor.DecompressFileToDirOptions).To(Equal([]boshcmd.CompressorOptions{{}}))
				decompressed = true
			}

			compressor.CompressFilesInDirCallBack = func() {
				Expect(fs.ReadFileString("/staging-release/license.tgz")).To(Equal("license-content"))
				compressed = true
			}

			path, err := act()
			Expect(err).ToNot(HaveOccurred())
			Expect(path).To(Equal("/release-archive"))

			Expect(compressed).To(BeTrue())
			Expect(decompressed).To(BeTrue())
		})

		It("returns error if copying license fails", func() {
			fs.CopyFileError = errors.New("fake-err")

			release.LicenseReturns(boshlic.NewLicense(
				NewResourceWithBuiltArchive("lic", "", "/tmp/lic.tgz", "")))

			_, err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))

			Expect(fs.FileExists("/staging-release")).To(BeFalse())
		})

		It("returns error if decompressing license fails", func() {
			fs.WriteFileString("/tmp/lic.tgz", "license-content")

			compressor.DecompressFileToDirErr = errors.New("fake-err")

			release.LicenseReturns(boshlic.NewLicense(
				NewResourceWithBuiltArchive("lic", "", "/tmp/lic.tgz", "")))

			_, err := act()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("fake-err"))

			Expect(fs.FileExists("/staging-release")).To(BeFalse())
		})
	})
})
