package resource

import (
	gopath "path"
	"sort"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	bicrypto "github.com/cloudfoundry/bosh-init/crypto"
)

type FingerprinterImpl struct {
	sha1calc bicrypto.SHA1Calculator
	fs       boshsys.FileSystem
}

func NewFingerprinterImpl(sha1calc bicrypto.SHA1Calculator, fs boshsys.FileSystem) FingerprinterImpl {
	return FingerprinterImpl{sha1calc: sha1calc, fs: fs}
}

func (f FingerprinterImpl) Calculate(files []File, additionalChunks []string) (string, error) {
	chunks := []string{"v2"}

	// Ensure consistent ordering of files
	sortedFiles := make([]File, len(files))
	copy(sortedFiles, files)
	sort.Sort(FileRelativePathSorting(sortedFiles))

	for _, file := range sortedFiles {
		chunk, err := f.fingerprintPath(file)
		if err != nil {
			return "", bosherr.WrapErrorf(err, "Fingerprinting file '%s'", file.Path)
		}

		chunks = append(chunks, chunk)
	}

	if len(additionalChunks) > 0 {
		// Ensure consistent ordering of additional chunks
		sortedAdditionalChunks := make([]string, len(additionalChunks))
		copy(sortedAdditionalChunks, additionalChunks)
		sort.Sort(AdditionalChunkSorting(sortedAdditionalChunks))

		chunks = append(chunks, strings.Join(sortedAdditionalChunks, ","))
	}

	return f.sha1calc.CalculateString(strings.Join(chunks, "")), nil
}

// fingerprintPath currently works with:
// - pkg: [rel_path, digest, is_hook ? '' : file_mode]
//   - changes: sorting
// - job: [File.basename(abs_path), digest, file_mode]
//   - changes: rel_path, sorting
// - lic: [File.basename(abs_path), digest]
//   - changes: sorting
func (f FingerprinterImpl) fingerprintPath(file File) (string, error) {
	var result string

	if file.UseBasename {
		result += gopath.Base(file.Path)
	} else {
		result += file.RelativePath
	}

	fileInfo, err := f.fs.Stat(file.Path)
	if err != nil {
		return "", err
	}

	if !fileInfo.IsDir() {
		sha1, err := f.sha1calc.Calculate(file.Path)
		if err != nil {
			return "", err
		}

		result += sha1
	}

	if !file.ExcludeMode {
		// Git doesn't really track file permissions, it just looks at executable
		// bit and uses 0755 if it's set or 0644 if not. We have to mimic that
		// behavior in the fingerprint calculation to avoid the situation where
		// seemingly clean working copy would trigger new fingerprints for
		// artifacts with changed permissions. Also we don't want current
		// fingerprints to change, hence the exact values below.
		var modeStr string

		if fileInfo.IsDir() {
			modeStr = "40755"
		} else if fileInfo.Mode()&0111 != 0 {
			modeStr = "100755"
		} else {
			modeStr = "100644"
		}

		result += modeStr
	}

	return result, nil
}

type AdditionalChunkSorting []string

func (s AdditionalChunkSorting) Len() int           { return len(s) }
func (s AdditionalChunkSorting) Less(i, j int) bool { return s[i] < s[j] }
func (s AdditionalChunkSorting) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
