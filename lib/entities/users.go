package entities

/*
  * Defines the different kind of ussers that can be able to use odin
*/

import (
  "owl/lib/utils"
)

type UserData struct {
  UserID string `json: "userid"`
  Phone string  `json: "phone"`
  UserName string `json: "username"`
  Email string  `json: "email"`
  Password string `json: "password"`
  Validated bool  `json: "anonymous"`
  Active bool `json: "verify"`
  utils.TimeStamps
}
