package classpath

import (
  "archive/zip"
  "errors"
  "io/ioutil"
  "path/filepath"
)

// 缓存打开的文件

type ZipEntry3 struct {
  absPath   string
  className []string
}

func (self *ZipEntry3) readClass(className string) ([]byte, Entry, error) {
  self.existsClassName(className)
  r, err := zip.OpenReader(self.absPath)
  if err != nil {
    return nil, nil, err
  }
  defer r.Close()
  for _, f := range r.File {
    if f.Name == className {
      rc, err := f.Open()
      if err != nil {
        return nil, nil, err
      }
      defer rc.Close()
      data, err := ioutil.ReadAll(rc)
      if err != nil {
        return nil, nil, err
      }
      return data, self, nil
    }
  }
  return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry3) existsClassName(className string) bool {
  if self.className == nil {
    self.calculate()
  }
  return false
}

func newZipEntry3(path string) *ZipEntry3 {
  absPath, err := filepath.Abs(path)
  if err != nil {
    panic(err)
  }
  return &ZipEntry3{absPath, nil}
}
func (self *ZipEntry3) String() string {
  return self.absPath
}

func (self *ZipEntry3) calculate() {
  //r, err := zip.OpenReader(self.absPath)

}
