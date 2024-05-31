package entities

import (
  "owl/lib/utils"
)

type IOC struct {
  IocID int
  VirusID string
  VType string
  Value string
  Source string
  utils.TimeStamps
}
