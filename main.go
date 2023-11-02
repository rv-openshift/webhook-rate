package main

import (	
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"net/smtp"
	"context"
	// "time"
	// "math/rand"

	"webhook/internal/rate"
	"webhook/internal/httprouter"
)

type payload_struct struct {               // json for argocd api payload
	AppName    string `json:"appName"`
	ArgoUrl    string `json:"argoUrl"`
	ArgoUser   string `json:"argoUser"`
	RepoPath   string `json:"repoPath"`
	RepoRef    string `json:"repoRef"`
	RepoUrl    string `json:"repoUrl"`
	Revision   string `json:"revision"`
	Sha        string `json:"sha"`
	Status     string `json:"status"`
	SyncStatus string `json:"syncStatus"`
	Timestamp  string `json:"timestamp"`
}

func SendmaiL(email, sync, url string, t payload_struct) bool {
	from := "venerayan@gmail.com"
	to := []string{
	  email,
	}
	b, ers := json.MarshalIndent(t, "", "  ")      // make pretty json
	if ers != nil {
		log.Println(ers)
		return false
	}
  	message := []byte("Subject: ["+sync+"]"+" Trigger of " + url + "\r\n" +
                    "\r\n" + 
					string(b) + "\r\n")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"	
	password := os.Getenv("emailpwd")		
	auth := smtp.PlainAuth("", email, password, smtpHost)
  	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
  	if err != nil {
		log.Println(err)
		return false
  	}
	return true
}

func endpointHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {  // endpoint of /
	body, e := io.ReadAll(request.Body)   // read the payload
	if e != nil {
		log.Println(e)
		return
	}
	var t payload_struct    // init variable for the json of argocd api payload; reference struct from mail package
	er := json.Unmarshal(body, &t)     // read the body to t as json readable
	if er != nil {
		// log.Println("No payload received!")
		return
	}
	log.Println("Payload: " + string(body))    // log the argocd api payload
	success := SendmaiL(t.ArgoUser, t.SyncStatus, t.RepoUrl, t) 
	if success {
		log.Println("Email Sent Successfully!")
	} else {
		log.Println("Email NOT Sent!")
	}

	
}

func rateLimiter(next func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)) httprouter.Handle {
	qps := 3
	burst := 1
	ctx := context.TODO()
	limit := rate.Limit(qps)
	limiter := rate.NewLimiter(limit, burst)	
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		for i := uint(1); i < 100; i++ {
			if !limiter.Allow() {
				if r.Method == "POST" {
					log.Println("API Blocked!")
				}
				_ = limiter.Wait(ctx)	
				return
			} else {
				// time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				next(w, r, nil)
			}
		}
	})
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    json.NewEncoder(w).Encode("OK")
}

func main() {
	log.Println("Webhook is running...")
	router := httprouter.New()
    router.GET("/", Index)  
    router.POST("/argocd-notify", rateLimiter(endpointHandler))  

    log.Fatal(http.ListenAndServe(":8080", router))
}
