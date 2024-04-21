package main

import "fmt"

// Yang dipilih untuk menerima pesan
type MessageSubscriber interface {
	Update(message string)
}

// Mengirim pesan yang terdaftar
type MessagePublisher struct {
	subscribers []MessageSubscriber
}

// Menambah penerima pesan
func (p *MessagePublisher) Attach(subscriber MessageSubscriber) {
	p.subscribers = append(p.subscribers, subscriber)
}

// mengirimkan pesan kepada semua penerima yang terdaftar.
func (p *MessagePublisher) Notify(message string) {
	for _, subscriber := range p.subscribers {
		subscriber.Update(message)
	}
}

// mewakili penerima pesan melalui email
type EmailSubscriber struct {
	email string
}

// metode yang dipanggil ketika EmailSubscriber menerima pesan baru
func (e *EmailSubscriber) Update(message string) {
	fmt.Printf("Email to %s: %s\n", e.email, message)
}

// mewakili penerima pesan melalui SMS
type SMSSubscriber struct {
	phoneNumber string
}

// metode yang dipanggil ketika SMSSubscriber menerima pesan baru
func (s *SMSSubscriber) Update(message string) {
	fmt.Printf("SMS to %s: %s\n", s.phoneNumber, message)
}

func main() {
	publisher := &MessagePublisher{}

	// Membuat penerima pesan
	emailSubscriber := &EmailSubscriber{email: "example@example.com"}
	smsSubscriber := &SMSSubscriber{phoneNumber: "+123456789"}

	// Menambahkan penerima pesan ke penerbit
	publisher.Attach(emailSubscriber)
	publisher.Attach(smsSubscriber)

	// Mengirimkan pesan
	publisher.Notify("New message: Hello, boy!")
}
