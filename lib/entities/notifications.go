package entities


import (
  "owl/lib/utils"
)

type Notifications struct {
  NotificationID string
  Title string
  Description string
  Read bool
  utils.TimeStamps
}
