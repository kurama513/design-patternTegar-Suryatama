package main

import (
  "bytes"
  "fmt"
  "testing"
)

// CustomWriter implements the io.Writer interface for testing purposes
type CustomWriter struct {
  content string
}

// Write method writes content to CustomWriter
func (w *CustomWriter) Write(p []byte) (n int, err error) {
  w.content = string(p)
  return len(p), nil
}

func TestMessagePublisher(t *testing.T) {
  // Create a custom writer to capture output
  capturedOutput := CustomWriter{}
  t.Run("TestNotify", func(t *testing.T) {
    // Redirect standard output to custom writer
    old := redirectStdout(&capturedOutput)
    defer resetStdout(old)

    publisher := &MessagePublisher{}

    // Create subscribers
    emailSubscriber := &EmailSubscriber{email: "example@example.com"}
    smsSubscriber := &SMSSubscriber{phoneNumber: "+123456789"}

    // Attach subscribers to publisher
    publisher.Attach(emailSubscriber)
    publisher.Attach(smsSubscriber)

    // Send message
    publisher.Notify("New message: Hello, World!")

    expectedOutput := "Email to example@example.com: New message: Hello, World!\nSMS to +123456789: New message: Hello, World!\n"

    // Check if the captured output matches the expected output
    if capturedOutput.content != expectedOutput {
      t.Errorf("Expected: %s\nGot: %s", expectedOutput, capturedOutput.content)
    }
  })
}

// redirectStdout redirects standard output and returns the original value
func redirectStdout(w *CustomWriter) *bytes.Buffer {
  old := stdout
  stdout = w
  return old
}

// resetStdout resets standard output to the original value
func resetStdout(old *bytes.Buffer) {
  stdout = old
}

var stdout = &CustomWriter{}

func main() {
  // Run tests
  fmt.Println("Running tests...")
  testing.Main(func(pat, str string) (bool, error) { return true, nil }, nil, nil, nil)
}
