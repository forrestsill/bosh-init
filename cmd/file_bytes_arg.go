package cmd

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type FileBytesArg struct {
	FS boshsys.FileSystem

	Bytes []byte
	Path  string
}

func (a *FileBytesArg) UnmarshalFlag(data string) error {
	if len(data) == 0 {
		return bosherr.Errorf("Expected file path to be non-empty")
	}

	absPath, err := a.FS.ExpandPath(data)
	if err != nil {
		return bosherr.WrapErrorf(err, "Getting absolute path '%s'", data)
	}

	bytes, err := a.FS.ReadFile(absPath)
	if err != nil {
		// todo go-flags does not show error when parsing
		return err
	}

	(*a).Bytes = bytes
	(*a).Path = absPath

	return nil
}
