package tourofgo

// Concurrency

	// Goroutines
	/*
		1. Definition - goroutine is a lightweighted thread managed by go runtime.
		2. go f(x,y,z) - starts a new goroutine running.
		3. f(x,y,z) - evaluation of f,x,y,z happens in same goroutine and the execution of f() happens in new goroutine.
		4. Goroutines run in same address space, so access to shared memory must be synchronized.
	*/
	// End of Goroutines

	// channels
	/*
		1. Definition - channnels are typed conduit through which you send and recieve values with the channel operator.
		2. <- channel operator - data flows in the direction of arrow.
		3. ch := make(chan int) - like map and slice you need to create channel before use.
		4. ch <- v - send v to channel ch.
		5. v := <-ch - recieve from channel ch and assign it to v.
		6. By default send and recieve block until the other side is ready. this allows goroutines to synchronize without explicit locks or condition variables.

		// BUffered Channels
		1. channels can be buffered, pass the buffered length as a second parameter to make() function to intialize a buffered channel.
		2. ch := make(chan int, 100).
		3. sends to buffer channel block only when the buffer is full, recieves block when the buffer is empty.
		// End of buffered channels

		// Range and Close
		1. sender can close a channel indicating that no more values will be sent.
		2. Reciever can test whether a channel has been closed by assigning a second parameter to the recieving expression.
		3. v, ok := <-ch - "ok is false" if there are no more values to be recieved and channel is closed.
		4. the loop for i:= range ch - will recieve values from the channel repeatedly until it is closed.
		5. only the sender can close the channel not the reciever, sending values to a closed channel will result in panic.
		6. channels are not like files, you need not to explicitly close them. closing only neccessary when the
		reciever must be told that no more values coming, such as terminating a range loop.
		// end of Range and Close.

		// select
		1. Def - select statement lets a goroutine wait on mutiplle communication operations.
		2. select blocks until its one the case can run, then it executes that case. It choose one at random if multiple are ready.
		3. Default case in select will run if no other case is ready. 
		// end of select
	*/
	// end of channels

	// Mutex
	/* 
	1. Definition - Mutex is a data structure that make sure only one goroutine can access a variable at a time to avoid conflicts.
	2. this concept is also called mutex exclusion.
	3. two methods of mutex standard library are - Lock, Unlock
	*/
	// End of Mutex

	// End of Concurrency
