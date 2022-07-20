package utils

import (
  "errors"
  //"fmt"
  "io/fs"
  //"log"
  "os"
  //"os/exec"
  "path/filepath"
  //"strings"
)

var Paths[] string

var walkDirFunc = func(path string, info fs.DirEntry, err error) error {
 if info.IsDir() {
   _ = filepath.SkipDir
 } else {
 Paths = append(Paths, path)
 }

  return nil
}

func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(err, os.ErrNotExist)
}

/* TODO: Implement building.
func Build(){
  if Exists("Makefile") {
    // Build
    cmd := exec.Command("make")
    if err := cmd.Run(); err != nil {fmt.Println(err)}
    // Instalation
    makecmd := exec.Command("make", "install")
    makecmd.Env = append(os.Environ(), "DESTDIR=work")
    out, _ := makecmd.Output()
    fmt.Println(string(out[:]))
    // Partial tracking
    filepath.WalkDir(".", walkDirFunc)
  }
}
*/
