package main

import (
  "flag"
  "fmt"
  "go-jvm/src/jvm/ch07/classpath"
  "go-jvm/src/jvm/ch07/rtda/heap"
  "os"
  "strings"
)

// java [-options] class [args...]
type Cmd struct {
  helpFlag         bool
  versionFlag      bool
  verboseClassFlag bool
  verboseInstFlag  bool
  cpOption         string
  XjreOption       string
  class            string
  args             []string
}

func ParseCmd() *Cmd {
  cmd := &Cmd{}

  flag.Usage = printUsage
  flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
  flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
  flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
  flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
  //flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
  flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose output")
  flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
  flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
  flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
  flag.Parse()

  args := flag.Args()
  if len(args) > 0 {
    cmd.class = args[0]
    cmd.args = args[1:]
  }

  return cmd
}

func printUsage() {
  fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func main() {
  cmd := ParseCmd()
  if cmd.versionFlag {
    fmt.Println("version 0.0.1")
  } else if cmd.helpFlag || cmd.class == "" {
    printUsage()
  } else {
    startJVM(cmd)
  }
}

func startJVM(cmd *Cmd) {
  cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
  classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)

  className := strings.Replace(cmd.class, ".", "/", -1)
  mainClass := classLoader.LoadClass(className)
  mainMethod := mainClass.GetMainMethod()
  if mainMethod != nil {
    interpret(mainMethod, cmd.verboseInstFlag)
  } else {
    fmt.Printf("Main method not found in class %s\n", cmd.class)
  }
}
