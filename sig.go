package fak

func SignalContext(ctx context.Context, signals ...os.Signal) context.Context {
	if len(signals) <= 0 {
		panic(errors.New("empty signals"))
	}
	var cancelGlobalContext context.CancelFunc
	globalContext, cancelGlobalContext := context.WithCancel(ctx)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		signalChan := make(chan os.Signal, len(signals)+1)
		defer fak.Flush(signalChan)
		defer cancelGlobalContext()
		signal.Notify(signalChan, signals...)
		wg.Done()
		select {
		case <-signalChan:
		case <-globalContext.Done():
		}
	}()
	wg.Wait()
	return globalContext
}
