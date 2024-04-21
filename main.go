// Observer pattern implementation in Go

package main

import "fmt"

// Observer interface
type MessageSubscriber interface {
  Update(message string)
}

// Subject struct
type MessagePublisher struct {
  subscribers []MessageSubscriber
}

// Attach method to add new subscribers
func (p *MessagePublisher) Attach(subscriber MessageSubscriber) {
  p.subscribers = append(p.subscribers, subscriber)
}

// Notify method to send message to all subscribers
func (p *MessagePublisher) Notify(message string) {
  for _, subscriber := range p.subscribers {
    subscriber.Update(message)
  }
}

// ConcreteObserver struct
type EmailSubscriber struct {
  email string
}

// Update method for EmailSubscriber
func (e *EmailSubscriber) Update(message string) {
  fmt.Printf("Email to %s: %s\n", e.email, message)
}

// ConcreteObserver struct
type SMSSubscriber struct {
  phoneNumber string
}

// Update method for SMSSubscriber
func (s *SMSSubscriber) Update(message string) {
  fmt.Printf("SMS to %s: %s\n", s.phoneNumber, message)
}

func main() {
  publisher := &MessagePublisher{}

  // Create subscribers
  emailSubscriber := &EmailSubscriber{email: "example@example.com"}
  smsSubscriber := &SMSSubscriber{phoneNumber: "+123456789"}

  // Attach subscribers to publisher
  publisher.Attach(emailSubscriber)
  publisher.Attach(smsSubscriber)

  // Send message
  publisher.Notify("New message: Hello, World!")
}
