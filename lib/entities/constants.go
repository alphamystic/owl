package entities

import (
  "errors"
  "database/sql"
)

var UserNotLoggedIn = errors.New("User Not Logged in.")

var NoCLaims = errors.New("No Claims")



const (
  ServerIssues = "We are experiencing internal server issues. Please try again later."
  NoCLaims = "No claims to return"
  UndefinedReconWebData = "Requested for an undefined web data type, use 1 for directories,2 for parameters and 3 for files."
  InstallDirectory = "/.odin/"
  bin = InstallDirectory + "bin/" // binary files
  PluginsDirectory = bin +"plugins/" // plugins directory
  MSDirectory = bin + "motherships/"  // fgenerated motherships storage
  MinionDirectory = bin + "minions/" // generated minions directory
  DownloadsDirectory = InstallDirectory + "downloads/"
  ScreenShotsDir = InstallDirectory + "screenshots/"
  TempDir = InstallDirectory + "temp/"
)

// user home directory (populated on start up)
var HomeDirectory string
