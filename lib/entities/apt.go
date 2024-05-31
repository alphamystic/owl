package entities

import (
  "owl/lib/utils"
)

type Apt struct {
  Name string `json:"apt_name"`
  Code int `json:"apt_code"`
  Aptid int `json:"aptid"`
  Description string `json:"description"`
  Active bool `json:"active"`
  utils.TimeStamps
}
