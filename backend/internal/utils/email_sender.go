// package utils

// import (
// 	"eCommerce/internal/config"
// 	"eCommerce/internal/models"
// 	"fmt"

// 	"gopkg.in/gomail.v2"
// )

// func OrderConfirmationMail(to string, order *models.Order) {
// 	conf := config.LoadEnv()
// 	smtpEmail := conf.SMTP_EMAIL
// 	smtpPassword := conf.SMTP_EMAIL_PASSWORD
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", smtpEmail)
// 	m.SetHeader("To", to)
// 	m.SetHeader("Subject", "Order Has Been Confirmed ")
// 	m.SetBody("text/html", "<h1>Hello!</h1><p>{{order}}</p>")

// 	d := gomail.NewDialer("smtp.gmail.com", 587, smtpEmail, smtpPassword)

// 	if err := d.DialAndSend(m); err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Email sent!")
// }

package utils

import (
	"eCommerce/internal/config"
	"eCommerce/internal/models"
	"fmt"

	"gopkg.in/gomail.v2"
)

func OrderConfirmationMail(to string, order *models.Order) error {
	conf := config.LoadEnv()
	smtpEmail := conf.SMTP_EMAIL
	smtpPassword := conf.SMTP_EMAIL_PASSWORD

	// Build the HTML Body dynamically
	htmlBody := fmt.Sprintf(`
		<h2>Order Confirmation</h2>
		<p>Hello ,</p>
		<p>Your order <strong>#%d</strong> has been confirmed.</p>

		<h3>Order Details:</h3>
		<ul>
			<li><strong>Order ID:</strong> %d</li>
			<li><strong>Total Amount:</strong> Rs.%0.2f</li>
			<li><strong>Status:</strong> %s</li>
		</ul>

		<p>Thank you for shopping with us!</p>
	`, order.ID, order.ID, order.TotalPrice, order.IsDelivered)

	m := gomail.NewMessage()
	m.SetHeader("From", smtpEmail)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Order Confirmation")
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, smtpEmail, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	fmt.Println("Email sent!")
	return nil
}
