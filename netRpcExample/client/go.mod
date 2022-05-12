module joe/client/calculator

go 1.17

require (
	github.com/myuser/calculator v0.0.0
)

replace (
	github.com/myuser/calculator => ../core
)