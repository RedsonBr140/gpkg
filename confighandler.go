package main

import (
  "github.com/IbrahimShahzad/gonfigure"
  "strconv"
  "errors"
  "fmt"
  "os"
)

const FILEDEFAULTERROR = "Error %s the config file. IF you need help, open a issue with the following:"

var (
  configPath string = "./gpkg.conf"
  SearchSize int    = 7
  Colors     bool
)

func initConfig() error {
  conf := gonfigure.InitialiseINIobj()
  conf = gonfigure.InsertSection(conf, "preferences")
  conf, err := gonfigure.WriteParameterToSection(conf, "preferences", "searchsize", "7")
  if err != nil { return err }
  conf, err = gonfigure.WriteParameterToSection(conf, "preferences", "colors", "true")
  if err != nil { return err }
  
  err = gonfigure.WriteINIFile(conf, configPath)
  return err
}

func failToGet(value string, defaultValue string){
  fmt.Println("Failed trying to get", value, "Using default", defaultValue)
}

func init(){
  if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
    fmt.Println("error: Config file not found, creating")

    err := initConfig()
    if err != nil {
      fmt.Printf("%s\nerr: %s",FILEDEFAULTERROR, "creating", err)
      os.Exit(1)
    }

    fmt.Println("Config file created.")
  }

  conf, err := gonfigure.LoadINI(configPath)
  if err != nil {
    fmt.Printf("%s\nerr: %s", FILEDEFAULTERROR, "openning", err)
  }
  cfgsearchsize, err := gonfigure.GetParameterValue(conf, "preferences", "searchsize")
  if err != nil {
    failToGet("searchsize.", "7")
  } else {
    SearchSize, err = strconv.Atoi(cfgsearchsize)
    if err != nil {
      failToGet("searchsize.", "7")
      SearchSize = 7
    }
  }
  
  cfgcolors, _ := gonfigure.GetParameterValue(conf, "preferences", "colors")
  Colors,    _ = strconv.ParseBool(cfgcolors)
}

