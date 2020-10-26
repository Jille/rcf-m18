# rcf-m18

[![GoDoc](https://godoc.org/github.com/Jille/rcf-m18/rcfm18?status.svg)](https://godoc.org/github.com/Jille/rcf-m18/rcfm18)

This is a Go library that allows you to interface with the [RCF M18](https://www.rcf.it/en_US/products/product-detail/m-18/291173).
The RCF M18 is a headless sound mixer. It normally is controlled by the MixRemote app.

This library provides an API to control the M18 without the MixRemote app. I personally use it to have my home automation control all sound in my place.

I have no affiliation with RCF. Everything here is based on reverse engineering the MixRemote App.

## gen

The `gen` folder contains the code to (re)generate the API. `source.txt` contains a DSL that defines the protocol. `parser.go` parses `source.txt` and generates `generated.go` (in the rcfm18 folder). Running `make(1)` inside the `gen` folder should reread `source.txt` and regenerate `generated.go`.

## rcfm18

This is the real library. It exports two types: `RCFM18Client` and `RCFM18Server` (create them with `Dial` and `NewServer`). Both of them try to shield you from the underlying protocol as much as possible, but the abstraction is somewhat thinand pull requests improving this are very welcome :)

The `RCFM18Client` exposes a bunch of methods that directly send a request to the device.

Both the `RCFM18Client` and `RCFM18Server` structs both expose a lot of callbacks that can be used by setting the function in the struct.

## Underlying protocol

Hopefully you don't need to know this when using this library.

Communication uses the [OSC protocol](http://opensoundcontrol.org/introduction-osc) over UDP. There is no retransmission/message acknowledging.

Messages have an address of the form /01/23/xy_456. Where 01 is the page, 23 the channel, xy the type and 456 the control. `source.txt` defines what each of these messages actually mean.

Each message contains additional parameters, they are either: a float32, a string or a list of floats. The meaning of these values is different per address.

## Call types

There's three types of communication messages:

- requests (e.g. asking the serial number of the device): A client initiates a request and will get a callback when answered.
- connset (e.g. enabling optimized meter sending for this connection): A client changes a setting related to this connection.
- set (e.g. changing the main volume): A client changes a setting and all other clients will get a message with this change.

Note that there is no strict relation between a request and a response. For one, because UDP is used, a reply might never arrive. Second, the device usually doesn't reply to changes made. Third, making some changes causes other changes to be implied. This is done by the device sending out multiple messages after receiving a single message.

## Contributions

Contributions are very welcome! Let me know if you need my help extending the parser/codegen.
