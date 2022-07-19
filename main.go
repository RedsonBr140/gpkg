package main

import (
  "fmt"
  "log"
  "os"
  "github.com/go-git/go-git/v5"
  "github.com/RedsonBr140/gpkg/utils"
)

func Install(){
  url, err := utils.GetRepoURL(InstallFlag[0], InstallFlag[1])
  if err != nil {
    fmt.Println(Color.Red + "Error:" + Color.Reset, "package not found.")
  }

  git.PlainClone("./root/var/cache/" + InstallFlag[0] + "_" + InstallFlag[1], false, &git.CloneOptions{
    URL:      url,
    Progress: os.Stdout,
  })
}

func SearchFunction(){
repos, err := utils.SearchRepos(SearchFlagVar)
  if err != nil {
    log.Fatal("Cannot get the repositories list, please check network connection.")
  }

  if len(repos.Repositories) == 0 {
    log.Fatal("Package not found.")
  }

  if len(repos.Repositories) < SearchSize {
    SearchSize = len(repos.Repositories)
  }

  for i := 0; i < SearchSize;i++ {
    fmt.Printf("\n%s%s%s",
    Color.Cyan, repos.Repositories[i].GetFullName(),
    Color.Reset)

    fmt.Printf("\n\t%s",
    utils.GetDescriptionText(repos.Repositories[i].GetDescription()))

    fmt.Printf("\n\t%sStars: %d %s\033[m",
    Color.Yellow,
    repos.Repositories[i].GetStargazersCount(),
    utils.GetForkText(repos.Repositories[i].GetFork()))
  }
  fmt.Println("\n :: Found", repos.GetTotal(), "matching packages.")
}

func main(){
  if !IsColorsEnable {
    Color.DisableColors()
  }

  if len(os.Args) < 2 {
    fmt.Println(Color.Bold + Color.Red + "error:" + Color.Reset, "no operation specified. (use -h for help)")
  }
  if SearchFlagVar != "" {
    SearchFunction()
  }
}

