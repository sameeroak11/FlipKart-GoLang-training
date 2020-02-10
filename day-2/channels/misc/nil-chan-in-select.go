func worker(i int, ch chan Work, quit chan struct{}) {
	for {
		select {
		case w := <-ch:
			if quit == nil {
				w.Refuse(); fmt.Println("worker", i, "refused", w)
				break
			}
			w.Do(); fmt.Println("worker", i, "processed", w)
		case <-quit:
			fmt.Println("worker", i, "quitting")
			quit = nil
		}
	}
}

func main() {
	ch, quit := make(chan Work), make(chan struct{})
	go makeWork(ch)
	for i := 0; i < 4; i++ { go worker(i, ch, quit) }
	time.Sleep(5 * time.Second)
	close(quit)
	time.Sleep(2 * time.Second)
}
