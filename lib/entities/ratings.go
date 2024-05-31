package entities

import (
  "owl/lib/utils"
)

type Ratings struct {
  PluginID string
  PluginName string
  Hash string // acts like the ID
  PreviousHashes []map[string]string
  Rates float32
  Average int // Total value as per last count
  utils.TimeStamps
}

type Rater struct {
  UserID string
  PluginID string
  Comment string
  utils.TimeStamps
}
