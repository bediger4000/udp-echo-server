# Quick and dirty UDP-based echo server and client in Go

Toy UDP server and client in Go. I did this to acquaint myself
with Go's standard package `net`, and specificlly the UDP-related
functions.

Because I learned as I wrote, here's a weird stylistic mismatch between how
server.go does UDP sockets, and how client.go does some "general" socket thing.
I did client2.go with just UDP functions, not at all general.

### Echo protocol

The server listens for UDP packets on a particular port number.
It blocks on `net.ReadFromUDP()` method call.

Should the server ever receive bytes, it prints out how many bytes
it received from where, then writes that same number of bytes back
to wherever it received them from.

The client creates a UDP connection to some IP (v4 or v6) or hostname
and a port number, based on command line information. Then it writes
bytes of a string, also from the command line, to the UDP connection.
It waits until some bytes come back to it, or an error occurs. Then it
exists.

Simple, and yet full of problems. No timeouts, no set number of bytes.
Client or server could hang forever waiting for a packet that never arrives.

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
