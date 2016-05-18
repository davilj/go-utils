package fileutils

import (
	//"strings"
  //s"fmt"
  "io/ioutil"
	"testing"
  "reflect"
  "os"
	"bytes"
  "path/filepath"
)

func TestCp(t *testing.T) {
  fromFileStr := "testData/testDataFile.txt"
  toFileStr := "testData/copyTestData.txt"
	Cp(fromFileStr, toFileStr)


  dataExpect, err := ioutil.ReadFile(fromFileStr)
  if err!=nil {
    t.Errorf("Could not read %s", fromFileStr)
  }

  dataTest, err := ioutil.ReadFile(toFileStr)
  if err!=nil {
    t.Errorf("Could not read %s", toFileStr)
  }

  if !reflect.DeepEqual(dataExpect,dataTest) {
    t.Error("Copy test failed")
  }

  err = os.Remove(toFileStr)
  if err!=nil {
    t.Error("Could not cleanup")
  }

}

func TestCpDir(t *testing.T) {
  fromFileStr := "testDataDir/one"
  toFileStr := "testDataDir/two"
	CpDir(fromFileStr, toFileStr)
  one:=[]byte(buildDirStr(fromFileStr))
  two:=bytes.Replace([]byte(buildDirStr(toFileStr)),[]byte("two"),[]byte("one"),-1)

  if !reflect.DeepEqual(one,two) {
    t.Error("Copy Dir test failed")
  }
}

func buildDirStr(start string) string {
  files, err := ioutil.ReadDir(start)
  if err != nil {
    panic(err)
  }
  tmp := "";
  for _,file := range(files) {
    newDir:= start + string(filepath.Separator) + file.Name()
    tmp = tmp + newDir
    info, err := os.Stat(newDir)
    if (err!=nil) {
      panic(err)
    }

    if info.IsDir() {
      tmp = tmp + "|" + buildDirStr(newDir)
    }
  }
  return tmp

}
