package main

import (
  "net/http"
  "github.com/mailgun/mailgun-go"
  "context"
  "time"
  "log"
  "fmt"
  // http.ResponseWriter, r *http.Request"log"
)


func index(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/home", 302)
}

func home(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "home", nil)
  return

}

func aboutUs(w http.ResponseWriter, r *http.Request){
  viewrender.HTML(w, http.StatusOK, "aboutUs", nil)
  return

}

func attorneys(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "attorneys", nil)
  return

}

func realEstate(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "realEstate", nil)
  return

}

func business(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "business", nil)
  return

}

func civilLitigation(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "civilLitigation", nil)
  return

}

func resources(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "resources", nil)
  return

}

func contactUs(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "contactUs", nil)
  return

}

func criminalLitigation(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "criminalLitigation", nil)
  return

}

func lawFirms(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "lawFirms", nil)
  return

}

func emailForm(w http.ResponseWriter, r *http.Request){

	err := r.ParseForm()
	if err != nil {
		//TODO:Handle Error
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	_, err = SendSimpleMessage("sandbox281cd85ad2934744bac165e60b510e89.mailgun.org", "e0a00e77d8fa281132f2d550b1836048-9ce9335e-de3475ec", name, email, phone, message)
	if err != nil{
		log.Println(err)
		return
	}

	viewrender.HTML(w, http.StatusOK, "contactUs", nil)

}



func SendSimpleMessage(domain string, apiKey string, name string, email string, phone string, msg string) (string, error) {
    mg := mailgun.NewMailgun(domain, apiKey)
    emailBody := fmt.Sprintf(emailTemp, name, email, phone, msg)
    m := mg.NewMessage(
        "Excited User <mailgun@sandbox281cd85ad2934744bac165e60b510e89.mailgun.org>",
        "You Have Been Contacted!",
        emailBody,
        "mrasnak3@gmail.com",
    )

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
    defer cancel()

    _, id, err := mg.Send(ctx, m)
    return id, err
}

/*
func testEmail(w http.ResponseWriter, r *http.Request){

	subject := "test email"
	emailTo := "mrasnak3@gmail.com"

	tags := []string{
		"testEmail",
	}

	err = mailer.InitializeTemplates(tags)
	if err != nil {
		log.Println(err)
		viewrender.Text(w, http.StatusInternalServerError, "There was an issue processing the security email, please try again. If error persists please contact support@pyaanalytics.com.")
		return
	}

	tag := "testEmail"

	htmlpayload, err := mailer.GetTemplateHtml(tag, nil)
	if err != nil {
		log.Println(err)
		viewrender.Text(w, http.StatusInternalServerError, "There was an issue processing the security email, please try again. If error persists please contact support@pyaanalytics.com.")
		return
	}

	textpayload, err := mailer.GetTemplateText(tag, nil)
	if err != nil {
		log.Println(err)
		viewrender.Text(w, http.StatusInternalServerError, "There was an issue processing the security email, please try again. If error persists please contact support@pyaanalytics.com.")
		return
	}

	msg := mailer.PrepareMailgunMessage(emailTo, subject, htmlpayload, textpayload)

	if err := mailer.SendMailgunMessage(msg); err != nil {
		log.Println(err)
		viewrender.Text(w, http.StatusInternalServerError, "There was an issue sending the security email, please try again. If error persists please contact support@pyaanalytics.com.")
		return
	}

	viewrender.Text(w, http.StatusOK, "An email was sent to the address on file, please follow the instructions within 15 minutes to reset your password.")
}*/
