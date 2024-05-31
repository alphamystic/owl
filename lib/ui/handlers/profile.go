package handlers

import(
  "fmt"
  "net/http"
  "owl/lib/utils"
)


func (hnd *Handler) Profile(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("profile","profile.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"profile",nil)
}


func (hnd *Handler) Updateprofile(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("update_profile","update_profile.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"update_profile",nil)
}

func (hnd *Handler) Securityprofile(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("security_profile","security_profile.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"security_profile",nil)
}



func (hnd *Handler) Notificationprofiles(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.GetATemplate("notification_settings","notification_settings.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"notification_settings",nil)
}
