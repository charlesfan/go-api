package tempfile

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var rand uint32
var randmu sync.Mutex

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextSuffix() string {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand = r
	randmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}

func TempFileWithSuffix(dir, prefix, suffix string) (f *os.File, err error) {
	if dir == "" {
		dir = os.TempDir()
	}

	nconflict := 0
	for i := 0; i < 10000; i++ {
		name := filepath.Join(dir, prefix+nextSuffix()+suffix)
		f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if os.IsExist(err) {
			if nconflict++; nconflict > 10 {
				rand = reseed()
			}
			continue
		}
		break
	}
	return
}

func CopyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
