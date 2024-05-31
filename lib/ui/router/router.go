package router

import(
  "os"
  "log"
  "fmt"
  "time"
  "context"
  "syscall"
  "net/http"
  "os/signal"
  "owl/lib/utils"
  ent"owl/lib/entities"
  "owl/lib/ui/handlers"
)

type Router struct {
  //Mux *http.ServeMux
  HTTPSvr *http.Server
  HTTPSSvr *http.Server
}

// should probably receive a server
func NewRouter(httpsSvr,httpSvr *http.Server) *Router {
  return &Router {
    //Mux: mux,
    HTTPSvr: httpSvr,
    HTTPSSvr: httpsSvr,
  }
}


func (rtr *Router) Run(reg bool){
  // create your db connection
  dbConfig := ent.IntitializeConnector("root","","localhost","owl")
  dbConn,err := ent.NewMySQLConnector(dbConfig)
  if err != nil {
    utils.Warning(fmt.Sprintf("Error connecting to the DB. \n[-]   ERROR: %s",err))
    return
  }

  // create a request logger
  rl := utils.NewRequestLogger("./.data/requests_logs/",066)
  // create shutdown channels
  ShutdownCh := make(chan bool)
  DoneCh := make(chan bool)
  //create  your handler
  hnd,err := handlers.NewHandler(dbConn, ShutdownCh, DoneCh,rl)
  if err != nil {
    utils.Logerror(err)
    return
  }

  fmt.Println("Registering routes.......")

  //http.HandleFunc("/",hnd.Home)
  http.HandleFunc("/login",hnd.Login)
  http.HandleFunc("/register",hnd.Register)
  http.HandleFunc("/forgot-password",hnd.Forgotpassword)
  http.HandleFunc("/confirm-password",hnd.Confirmpassword)
  http.HandleFunc("/blank",hnd.Blank)
  http.HandleFunc("/dash",hnd.Dash)
  http.HandleFunc("/",hnd.Notfound)
  /* Client Routes */
  http.HandleFunc("/ms",hnd.Motherships)
  http.HandleFunc("/profile",hnd.Profile)
  http.HandleFunc("/updateprofile",hnd.Updateprofile)
  http.HandleFunc("/securityprofile",hnd.Securityprofile)
  http.HandleFunc("/notificationprofile",hnd.Notificationprofiles)
  http.HandleFunc("/listapikeys",hnd.ClientListapikeys)
  http.HandleFunc("/createapikeys",hnd.CreateApikeys)
  http.HandleFunc("/issues",hnd.Crm_Issues)
  http.HandleFunc("/appointments",hnd.Client_Appointments)
  http.HandleFunc("/appointment-detail",hnd.Client_Appointments_Detail)
  /// Change any template with index.html

  fmt.Println("Handlers are registered............")

  // create a file server for the static files
  fs := http.FileServer(http.Dir("./lib/ui/static"))
  // Cache static files for 1 hour (adjust as needed)
  http.Handle("/static/", http.StripPrefix("/static", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Cache-Control", "max-age=3600")
    fs.ServeHTTP(res,req)
  })))

  // create a file server for the downloadable files
  downloads_dir := http.FileServer(http.Dir("./lib/ui/downloads"))
  http.Handle("/downloads/", http.StripPrefix("/downloads", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Cache-Control", "max-age=3600")
    downloads_dir.ServeHTTP(res,req)
  })))

  // create a file server for the uploaded files
  uploads := http.FileServer(http.Dir("./lib/ui/uploads"))
  http.Handle("/uploads/", http.StripPrefix("/uploads", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Cache-Control", "max-age=3600")
    uploads.ServeHTTP(res,req)
  })))

  // Start the server on the background
   go func(){
     if err := rtr.HTTPSvr.ListenAndServe(); err != http.ErrServerClosed {
       log.Fatalf("[-] Error starting server: %s\n",err.Error())
     }
   }()
   go func(){
     // we need to find a better way of supplying this
     if err := rtr.HTTPSSvr.ListenAndServeTLS("./certs/server.crt", "./certs/server.key"); err != http.ErrServerClosed {
       log.Fatalf("[-] Error starting HTTPS server: %s\n",err.Error())
     }
   }()
   interruptChan := make(chan os.Signal,1)
   signal.Notify(interruptChan,os.Interrupt, syscall.SIGTERM)
   //sedn a close channel to the handler
   //hnd.ShutdownChan <- true
   // wait for the receiver to finish writing all logs
  // <-hnd.DoneChan
   // read from the interrupt chan and shutdown
   <-interruptChan
   shutdownCtx,shutdownCancel := context.WithTimeout(context.Background(),5 * time.Second)
   defer shutdownCancel()
   var errs []error
   if err = rtr.HTTPSvr.Shutdown(shutdownCtx); err != nil {
     log.Println("[-] Server shutdown error: %s\n",err.Error())
     errs = append(errs,err)
   }
   err = rtr.HTTPSSvr.Shutdown(shutdownCtx)
   if err != nil {
     log.Println("[-] Server shutdown error: %s\n",err.Error())
     errs = append(errs,err)
   }
   for _, err := range errs {
        log.Println(err)
    }
   log.Println("[+] Server gracefully stopped.")
}
