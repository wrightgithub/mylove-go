module helloxyy.com/mylove-go

go 1.17

replace (
	example.com/greetings => ../greetings
	helloxyy.com/utilsmod => ./utilsmod
)

require (
	example.com/greetings v1.1.0
	helloxyy.com/utilsmod v0.0.0-00010101000000-000000000000
)
