# curlbin, the curl-oriented pastebin

WARNING: project in early alpha stage

curlbin intended to be used by curl or other shell http clients. It is simple:
just pipe your text/logs/code chunks into `curl --data-binary @- $URL` and it
will give you link back. Curl it and it will give you content back, possibly
coloring the code or adding line numbers (see options below).

## Installation

`go get github.com/mechmind/curlbin`
For now, requires [pygmentize](http://pygments.org/docs/cmdline/) to work.

## Running

curlbin has those configuration options:

 * -datadir - directory for storing pastes;
 * -listen - address and port to listen on;
 * -logfile - log to given file. Set to "-" to use stdout
 * -server-name - server name with port that will be used in urls when adding
  pastes. If not specified, will reuse "Host" header from request.

## Usage

### Uploading pastes
Pipe some data into `curl --data-binary @- $URL`. Server will return an url for
that paste.

### Viewing paste
Curl url from previous phase to retrieve paste content, unchanged or filtered
(see below).

### View filters
TODO.
