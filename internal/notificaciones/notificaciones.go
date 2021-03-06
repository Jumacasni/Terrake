package notificaciones

import (
	"fmt"
	"net/smtp"
)

// Emails registrados en la base de datos, de momento solo es un map
// test@test.com = true -> el email tiene activadas las notificaciones
// test@test.com = false -> el email no tiene activadas las notificaciones
type Notificaciones struct {
	emails map[string]bool
}

var (
	db Notificaciones
)

// Inicializa la db con un email
func init() {
	InitEmail("prueba@prueba.com")
}

func InitEmail(email string){
	db.emails = make(map[string]bool)
	db.emails[email] = true
}

// Comprueba si el email existe en la db
func EmailExists(email string) error{
	_, exists := db.emails[email]
	if !exists{
		return fmt.Errorf("El email '%s' no existe", email)
	}
	return nil
}

// Comprueba si el email tiene activadas las notificaciones
func EmailActivated(email string) error{
	value, _ := db.emails[email]
	if !value{
		return fmt.Errorf("El email '%s' no tiene activadas las notificaciones", email)
	}
	return nil
}

// Mock para la función EmailExists
type EmailPreCheck interface{
	emailExists(string) bool
}

/****** AÑADIR EMAIL ******/
func AddEmail(email string, checkEmail EmailPreCheck) error{
	// Usa la función mock
	found := checkEmail.emailExists(email)
	if found{
		return fmt.Errorf("El email '%s' ya está registrado", email)
	}

	// Inserta el email en la db
	db.emails[email] = true
	return nil
}

/****** MODIFICA EMAIL ******/
func ModifyEmail(email string, newEmail string, checkEmail EmailPreCheck) error{
	// Usa la función mock
	found := checkEmail.emailExists(email)
	if !found{
		return fmt.Errorf("El email '%s' no existe", email)
	}

	// Se añade el nuevo email con el valor que tenía antes
	db.emails[newEmail] = db.emails[email]
	// Se borra el antiguo email
	delete(db.emails, email)

	return nil
}

// Mock para la función EmailActivated
type EmailActivatedPreCheck interface{
	emailActivated(string) bool
}

/****** DESACTIVA NOTIFICACIONES ******/
func DeactivateEmail(email string, checkEmail EmailPreCheck, checkEmailActivated EmailActivatedPreCheck) error{
	// Usa la función mock
	found := checkEmail.emailExists(email)
	if !found{
		return fmt.Errorf("El email '%s' no existe", email)
	}

	activated := checkEmailActivated.emailActivated(email)
	if !activated{
		return fmt.Errorf("El email '%s' ya tiene las notificaciones desactivadas", email)
	}

	// Desactiva las notificaciones
	db.emails[email] = false

	return nil
}

// Intefaz para mockear la función de enviar un email
type EmailSender interface {
	SendEmail(email []string, body []byte) error
}

// Implementación para la interfaz EmailSender, el servidor del email
type EmailConfig struct {
	user string
	password string
	host string
	port string
	email string
}

// Describe el struct que manda el email
type emailSender struct {
	conf EmailConfig
	send func(string, smtp.Auth, string, []string, []byte) error
}

// Simula la función smtp.SendMail
func NewEmailSender(conf EmailConfig) EmailSender {
	return &emailSender{conf, smtp.SendMail}
}

/***** SEND EMAIL *****/
func (e *emailSender) SendEmail(email []string, terremotos []byte) error{
	addr := e.conf.host + ":" + e.conf.port
	auth := smtp.PlainAuth("", e.conf.user, e.conf.password, e.conf.host)
	return e.send(addr, auth, e.conf.email, email, terremotos)
}