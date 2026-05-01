package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup){
	defer wg.Done()
	for recipient := range ch {
		smtpHost:= "localhost"
		smtpPort:= "1025"
		// formattedMsg := fmt.Sprintf("To: %s\r\rSubject: Test Email\r\n\r\n%s\r\n", recipient.Email, "test")
		// msg := []byte(formattedMsg)

		msg,err := executeTemplate(recipient)
		if err != nil{
			fmt.Printf("worker : %d error parsing templte for %s",id,recipient.Email)
			continue
		}

		fmt.Println(id,recipient.Email,recipient.Name)


		err = smtp.SendMail(smtpHost + ":" + smtpPort, nil, "avi@gmail.com", []string{recipient.Email}, []byte(msg))
		if err != nil{
			log.Fatal(err)
		}
		time.Sleep(50 * time.Millisecond)
		fmt.Println(id,recipient.Email,recipient.Name)
		
	}
}