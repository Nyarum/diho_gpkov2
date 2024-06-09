module github.com/Nyarum/diho_gpkov2

go 1.22.2

require (
	github.com/google/uuid v1.6.0
	github.com/valyala/bytebufferpool v1.0.0
)

require (
	github.com/Nyarum/diho_bytes_generate v0.0.6
	github.com/davecgh/go-spew v1.1.1
	go.etcd.io/bbolt v1.3.10
)

require (
	github.com/maruel/panicparse/v2 v2.3.1 // indirect
	golang.org/x/sys v0.4.0 // indirect
)

replace github.com/Nyarum/diho_bytes_generate => ../diho_bytes_generate
