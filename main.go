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
)

func main() {
	relPath := os.Args[1]

	path, err := filepath.Abs(relPath)
	errorOut(err)

	fi, err := os.Lstat(path)
	errorOut(err)

	h, err := archive.FileInfoHeader(path, relPath, fi)
	errorOut(err)

	tsh := &tarsumHash{hdr: h, Hash: sha256.New()}
	tsh.Reset()

	if fi.Mode().IsRegular() && fi.Size() > 0 {
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
		{"mtime", strconv.FormatInt(h.ModTime.UTC().Unix(), 10)},
		{"typeflag", string([]byte{h.Typeflag})},
		{"linkname", h.Linkname},
		{"uname", h.Uname},
		{"gname", h.Gname},
		{"devmajor", strconv.FormatInt(h.Devmajor, 10)},
		{"devminor", strconv.FormatInt(h.Devminor, 10)},
	}

	var headers [][2]string = make([][2]string, 0);
	headers = append(headers, orderedHeaders[0:5]...)
	headers = append(headers, orderedHeaders[6:]...)

	fmt.Printf("Name: %s\n", relPath)
	fmt.Printf("Hash: %s\n", hex.EncodeToString(tsh.Hash.Sum(nil)))
	fmt.Printf("Hdrs: %s\n", headers)
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