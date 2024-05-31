package entities

import (
  "owl/lib/utils"
)

type Api struct {
  ApiKey string `json:"apikey"`
  OwnerID string `json:"ownerid"` // refers to apikey owner
  Active bool `json:"active"`
  Custom bool `json:"custom"`
  utils.TimeStamps
}
