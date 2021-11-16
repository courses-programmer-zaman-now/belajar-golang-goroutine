package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// membuat channel
	channel := make(chan string)
	defer close(channel)

	// membuat goroutine anonymouse function
	go func() {
		time.Sleep(2 * time.Second)
		// mengirim data ke channel
		channel <- "Danil Syah"
		fmt.Println("Selesai Mengirim Data ke Channel")

	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// channel sebagai parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "haykal dafiansyah"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// data diterima oleh channel
	go GiveMeResponse(channel)
	fmt.Println("selesai mengirim data ke channel")

	// data channel di terima oleh variabel nama
	nama := <-channel
	fmt.Println(nama)

	time.Sleep(5 * time.Second)

}
