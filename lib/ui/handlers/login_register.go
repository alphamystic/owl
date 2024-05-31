package handlers


import(
  "fmt"
  "net/http"
  "owl/lib/utils"
)

func (hnd *Handler) Login(res http.ResponseWriter, req *http.Request){
  if req.Method == "POST" {
    req.ParseForm()
    pass := req.FormValue("password")
    email := req.FormValue("email")
    isAuth,presId := hnd.Authenticate(pass,email)
    if !isAuth {
      hnd.Tpl.ExecuteTemplate(res,"login.html","Wrong username or password")
      return
    }
    //set session
    session,_ := hnd.Store.Get(req,"session")
    session.Values["UserID"] = presId
    session.Save(req,res)
    //redirect to dashboard or get the dash data and execute dash
    http.Redirect(res,req,"/dash",http.StatusSeeOther)
    return
  }
  tpl,err := hnd.LoadATemplate("login","login.html")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"login",nil)
  return
}

func (hnd *Handler) Logout(res http.ResponseWriter, req *http.Request){
  session,_ := hnd.Store.Get(req,"session")
  delete(session.Values,"UserID")
  session.Save(req,res)
  tpl,err := hnd.LoadATemplate("blank","blank.tmpl")
  if err != nil{
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"login.html","Logged Out ADIOS!!!")
  return
}


func (hnd *Handler) Register(res http.ResponseWriter, req *http.Request){
  // check if user is logged in first
  if req.Method != "POST" {
    tpl,err := hnd.LoadATemplate("register","register.html")
    if err != nil {
      utils.Warning(fmt.Sprintf("%s",err))
      http.Error(res, "An error occurred", http.StatusInternalServerError)
    }
    tpl.ExecuteTemplate(res,"register",nil)
    return
  }
  tpl,err := hnd.LoadATemplate("register","register.html")
  if err != nil {
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"register",nil)
  return
}

func (hnd *Handler) Forgotpassword(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.LoadATemplate("forgot_password","forgot_password.html")
  if err != nil {
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"forgot_password",nil)
  return
}

func (hnd *Handler) Confirmpassword(res http.ResponseWriter, req *http.Request) {
  tpl,err := hnd.LoadATemplate("confirm_password","confirm_password.html")
  if err != nil {
    utils.Warning(fmt.Sprintf("%s",err))
    http.Error(res, "An error occurred", http.StatusInternalServerError)
  }
  tpl.ExecuteTemplate(res,"confirm_password","Password has been reset.")
  return
}
