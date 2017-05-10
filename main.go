package main

import (
	"os"
	"archive/tar"

	"fmt"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/tarsum"
	"github.com/docker/docker/pkg/pools"
	"path/filepath"
	"crypto/sha256"
	"hash"
	"encoding/hex"
	"strconv"
	"github.com/docker/docker/builder/remotecontext"
)

func main() {
	rootDir, err := filepath.Abs(".")
	errorOut(err)

	relPath := os.Args[1]

	path, err := filepath.Abs(relPath)
	errorOut(err)

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		rel, err := remotecontext.Rel(rootDir, path)
		errorOut(err)

		h, err := archive.FileInfoHeader(path, rel, info)
		errorOut(err)

		tsh := &tarsumHash{hdr: h, Hash: sha256.New()}
		tsh.Reset()

		if info.Mode().IsRegular() && info.Size() > 0 {
			//fileHash = hashFile(path, h)
			f, err := os.Open(path)
			errorOut(err)
			defer f.Close()
			if _, err := pools.Copy(tsh, f); err != nil {
				errorOut(err)
			}
		}

		orderedHeaders := [][2]string{
			{"name", h.Name},
			{"mode", strconv.FormatInt(h.Mode, 10)},
			{"uid", strconv.Itoa(h.Uid)},
			{"gid", strconv.Itoa(h.Gid)},
			{"size", strconv.FormatInt(h.Size, 10)},
			// {"mtime", strconv.FormatInt(h.ModTime.UTC().Unix(), 10)},
			{"typeflag", string([]byte{h.Typeflag})},
			{"linkname", h.Linkname},
			{"uname", h.Uname},
			{"gname", h.Gname},
			{"devmajor", strconv.FormatInt(h.Devmajor, 10)},
			{"devminor", strconv.FormatInt(h.Devminor, 10)},
		}

		fmt.Printf("Hash: %s, Hdrs: %s\n", hex.EncodeToString(tsh.Sum(nil)), orderedHeaders)

		return nil
	})
}

func errorOut(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

type tarsumHash struct {
	hash.Hash
	hdr *tar.Header
}

// Reset resets the Hash to its initial state.
func (tsh *tarsumHash) Reset() {
	// comply with hash.Hash and reset to the state hash had before any writes
	tsh.Hash.Reset()
	tarsum.WriteV1Header(tsh.hdr, tsh.Hash)
}