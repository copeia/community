// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

// jshint ignore:start

package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/documize/community/core/api/request"
	"github.com/documize/community/core/log"
	"github.com/documize/community/server/web"
)

// InviteNewUser invites someone new providing credentials, explaining the product and stating who is inviting them.
func InviteNewUser(recipient, inviter, url, username, password string) {
	method := "InviteNewUser"

	file, err := web.ReadFile("mail/invite-new-user.html")

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to load email template", method), err)
		return
	}

	emailTemplate := string(file)

	// check inviter name
	if inviter == "Hello You" || len(inviter) == 0 {
		inviter = "Your colleague"
	}

	subject := fmt.Sprintf("%s has invited you to Documize", inviter)

	e := NewEmail()
	e.From = SMTPCreds.SMTPsender()
	e.To = []string{recipient}
	e.Subject = subject

	parameters := struct {
		Subject  string
		Inviter  string
		Url      string
		Username string
		Password string
	}{
		subject,
		inviter,
		url,
		recipient,
		password,
	}

	buffer := new(bytes.Buffer)
	t := template.Must(template.New("emailTemplate").Parse(emailTemplate))
	log.IfErr(t.Execute(buffer, &parameters))
	e.HTML = buffer.Bytes()

	err = e.Send(GetHost(), GetAuth())

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to send email", method), err)
	}
}

// InviteExistingUser invites a known user to an organization.
func InviteExistingUser(recipient, inviter, url string) {
	method := "InviteExistingUser"

	file, err := web.ReadFile("mail/invite-existing-user.html")

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to load email template", method), err)
		return
	}

	emailTemplate := string(file)

	// check inviter name
	if inviter == "Hello You" || len(inviter) == 0 {
		inviter = "Your colleague"
	}

	subject := fmt.Sprintf("%s has invited you to their Documize account", inviter)

	e := NewEmail()
	e.From = SMTPCreds.SMTPsender()
	e.To = []string{recipient}
	e.Subject = subject

	parameters := struct {
		Subject string
		Inviter string
		Url     string
	}{
		subject,
		inviter,
		url,
	}

	buffer := new(bytes.Buffer)
	t := template.Must(template.New("emailTemplate").Parse(emailTemplate))
	log.IfErr(t.Execute(buffer, &parameters))
	e.HTML = buffer.Bytes()

	err = e.Send(GetHost(), GetAuth())

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to send email", method), err)
	}
}

// PasswordReset sends a reset email with an embedded token.
func PasswordReset(recipient, url string) {
	method := "PasswordReset"

	file, err := web.ReadFile("mail/password-reset.html")

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to load email template", method), err)
		return
	}

	emailTemplate := string(file)

	subject := "Documize password reset request"

	e := NewEmail()
	e.From = SMTPCreds.SMTPsender() //e.g. "Documize <hello@documize.com>"
	e.To = []string{recipient}
	e.Subject = subject

	parameters := struct {
		Subject string
		Url     string
	}{
		subject,
		url,
	}

	buffer := new(bytes.Buffer)
	t := template.Must(template.New("emailTemplate").Parse(emailTemplate))
	log.IfErr(t.Execute(buffer, &parameters))
	e.HTML = buffer.Bytes()

	err = e.Send(GetHost(), GetAuth())

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to send email", method), err)
	}
}

// ShareFolderExistingUser provides an existing user with a link to a newly shared folder.
func ShareFolderExistingUser(recipient, inviter, url, folder, intro string) {
	method := "ShareFolderExistingUser"

	file, err := web.ReadFile("mail/share-folder-existing-user.html")

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to load email template", method), err)
		return
	}

	emailTemplate := string(file)

	// check inviter name
	if inviter == "Hello You" || len(inviter) == 0 {
		inviter = "Your colleague"
	}

	subject := fmt.Sprintf("%s has shared %s with you", inviter, folder)

	e := NewEmail()
	e.From = SMTPCreds.SMTPsender()
	e.To = []string{recipient}
	e.Subject = subject

	parameters := struct {
		Subject string
		Inviter string
		Url     string
		Folder  string
		Intro   string
	}{
		subject,
		inviter,
		url,
		folder,
		intro,
	}

	buffer := new(bytes.Buffer)
	t := template.Must(template.New("emailTemplate").Parse(emailTemplate))
	log.IfErr(t.Execute(buffer, &parameters))
	e.HTML = buffer.Bytes()

	err = e.Send(GetHost(), GetAuth())

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to send email", method), err)
	}
}

// ShareFolderNewUser invites new user providing credentials, explaining the product and stating who is inviting them.
func ShareFolderNewUser(recipient, inviter, url, folder, invitationMessage string) {
	method := "ShareFolderNewUser"

	file, err := web.ReadFile("mail/share-folder-new-user.html")

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to load email template", method), err)
		return
	}

	emailTemplate := string(file)

	// check inviter name
	if inviter == "Hello You" || len(inviter) == 0 {
		inviter = "Your colleague"
	}

	subject := fmt.Sprintf("%s has shared %s with you on Documize", inviter, folder)

	e := NewEmail()
	e.From = SMTPCreds.SMTPsender()
	e.To = []string{recipient}
	e.Subject = subject

	parameters := struct {
		Subject    string
		Inviter    string
		Url        string
		Invitation string
		Folder     string
	}{
		subject,
		inviter,
		url,
		invitationMessage,
		folder,
	}

	buffer := new(bytes.Buffer)
	t := template.Must(template.New("emailTemplate").Parse(emailTemplate))
	log.IfErr(t.Execute(buffer, &parameters))
	e.HTML = buffer.Bytes()

	err = e.Send(GetHost(), GetAuth())

	if err != nil {
		log.Error(fmt.Sprintf("%s - unable to send email", method), err)
	}
}

// SMTPCreds return SMTP configuration.
var SMTPCreds = struct{ SMTPuserid, SMTPpassword, SMTPhost, SMTPport, SMTPsender func() string }{
	func() string { return request.ConfigString("SMTP", "userid") },
	func() string { return request.ConfigString("SMTP", "password") },
	func() string { return request.ConfigString("SMTP", "host") },
	func() string {
		r := request.ConfigString("SMTP", "port")
		if r == "" {
			return "587" // default port number
		}
		return r
	},
	func() string { return request.ConfigString("SMTP", "sender") },
}

// GetAuth to return SMTP credentials
func GetAuth() smtp.Auth {
	a := smtp.PlainAuth("", SMTPCreds.SMTPuserid(), SMTPCreds.SMTPpassword(), SMTPCreds.SMTPhost())
	//fmt.Printf("DEBUG GetAuth() = %#v\n", a)
	return a
}

// GetHost to return SMTP host details
func GetHost() string {
	h := SMTPCreds.SMTPhost() + ":" + SMTPCreds.SMTPport()
	//fmt.Printf("DEBUG GetHost() = %#v\n", h)
	return h
}
