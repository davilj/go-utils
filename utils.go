package fileutils

import (
  "os"
  "io"
  "io/ioutil"
  //"fmt"
  "path/filepath"
)

func Cp(fromFileStr string, toFileStr string) {
  fromFile,err := os.Open(fromFileStr)
  if err!=nil {
    panic("Could not open [" + fromFileStr + "]")
  }
  defer fromFile.Close()

  toFile,err := os.Create(toFileStr)
  if err!=nil {
    panic("Could not open [" + toFileStr + "]")
  }
  defer toFile.Close()

  _, err = io.Copy(toFile, fromFile)
  if err!=nil {
    panic("Could not copy file")
  }

  err = toFile.Sync()
  if err!=nil {
    panic("Could not copy file")
  }
}

func CpDir(fromDir string, toDir string) {
  files, err := ioutil.ReadDir(fromDir)
  if err != nil {
    panic("Could not read source dir ")
  }

  for _,file := range(files) {
    newFromDir := fromDir + string(filepath.Separator) + file.Name()
    newToDir := toDir + string(filepath.Separator) + file.Name()
    //dir := filepath.Dir(newToDir)
    os.MkdirAll(toDir, 0777)
    info, err := os.Stat(newFromDir)
    if (err!=nil) {
      panic(err)
    }

    if info.IsDir() {
      CpDir(newFromDir, newToDir)
    } else {
      Cp(newFromDir, newToDir)
    }
  }

}
