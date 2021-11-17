package belajargolanggoroutine

import (
	"fmt"
	"strconv"
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

// channel In dan Out

// hanya mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Danil Syah"
}

// hanya menerima data chanel ke variabel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// belajar buffered channel
func TestBufferedChannel(t *testing.T) {
	// buffered : kapasitas menampung data dalam satu channel
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "danil"
	channel <- "syah"
	channel <- "ari"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	// fmt.Println(<-channel)

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	fmt.Println("Selesai")
}

// belajar range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data ", data)
	}

	fmt.Println("Selesai")
}

// belajar Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Dataa dari channel 2 ", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}
