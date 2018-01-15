package cp

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func replaceHomeFolder(path string) (string, error) {
	if !strings.HasPrefix(path, "~") {
		return path, nil
	}
	var buffer bytes.Buffer
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	_, err = buffer.WriteString(usr.HomeDir)
	if err != nil {
		return "", err
	}
	_, err = buffer.WriteString(strings.TrimPrefix(path, "~"))
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func AbsolutePath(path string) (string, error) {
	homeReplaced, err := replaceHomeFolder(path)
	if err != nil {
		return "", err
	}
	return filepath.Abs(homeReplaced)
}

// CopyFile copies a file from src to dst. If src and dst files exist, and are
// the same, then return success. Otherwise, attempt to create a hard link
// between the two files. If that fails, copy the file contents from src to dst.
func CopyFile(src, dst string) (err error) {
	srcAbs, err := AbsolutePath(src)
	if err != nil {
		return err
	}
	dstAbs, err := AbsolutePath(dst)
	if err != nil {
		return err
	}

	// open source file
	sfi, err := os.Stat(srcAbs)
	if err != nil {
		return
	}
	if !sfi.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	// open dest file
	dfi, err := os.Stat(dstAbs)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
		// file doesn't exist
		spew.Dump(filepath.Dir(dst))
		err := os.MkdirAll(filepath.Dir(dst), 0755)
		if err != nil {
			spew.Dump(err)
		}

	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}
		if os.SameFile(sfi, dfi) {
			return
		}
	}
	if err = os.Link(src, dst); err == nil {
		return
	}
	return copyFileContents(src, dst)
}

// copyFileContents copies the contents of the file named src to the file named
// by dst. The file will be created if it does not already exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file.
func copyFileContents(src, dst string) (err error) {
	// Open the source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	// Open the destination file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return
	}
	// Return any errors that result from closing the destination file
	// Will return nil if no errors occurred
	defer func() {
		cerr := dstFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Copy the contents of the source file into the destination files
	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return
	}
	err = dstFile.Sync()
	return
}
