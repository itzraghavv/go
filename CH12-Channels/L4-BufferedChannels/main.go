package main

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))

	done := make(chan struct{})

	go func() {
		for _, email := range emails {
			ch <- email
		}
		close(done)
	}()

	<-done

	return ch
}
