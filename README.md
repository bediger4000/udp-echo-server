# Quick and dirty UDP-based echo server and client in Go

There's a weird stylistic mismatch between how server.go does UDP
sockets, and how client.go does some "general" socket thing.
I did client2.go with just UDP functions, not at all general.

## Building

    $ go build server.go
    $ go build client.go
    $ go build client2.go

## Usage

In Window 1:

    $ ./server :: 7890
    Accepting a new packet


In Window 2:

    $ ./client udp localhost 7890 'some string'

Or:

    $ ./client2 fe80::a11:96ff:fe7f:6d74 7890 'some string' [eth0]

Note the contents of `net.UDPAddr`:

    type UDPAddr struct {
            IP   IP
            Port int
            Zone string // IPv6 scoped addressing zone
    }

The Zone element is used for routing link-local addresses
(look like fe80::), the interface name works as the zone.
Called with a 4th argument, `client2` uses that argument
as the "zone".
