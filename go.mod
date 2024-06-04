module github.com/Nyarum/diho_gpkov2

go 1.22.2

require (
	github.com/google/uuid v1.6.0
	github.com/valyala/bytebufferpool v1.0.0
)

require github.com/Nyarum/diho_bytes_generate v0.0.6

replace github.com/Nyarum/diho_bytes_generate => ../diho_bytes_generate
