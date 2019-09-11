# truspyd - A multi-platform service for tracking TCP, RAW and UDP sockets

[![Travis](https://travis-ci.org/shuntingyard/truspyd.svg?branch=master)](https://travis-ci.org/shuntingyard/truspyd)
[![License](https://github/license/shuntingyard/truspyd)](http://github.com/shuntingyard/flowproc/blob/master/LICENSE.txt)

Track socket details in order to gain network activity insights on a
per process/program level and make these insights available to the analytics
modules of a running [flowproc](https://github.com/shuntingyard/flowproc)
instance.

(TRU stands for TCP, RAW and UDP of course - what else? :)

## Implementation Notes

### Technical

Intended to run as daemon/service on supported platforms (Linux, Windows
and OS X).

The current implementation calls existing utilities like ``ss`` or ``netstat``
periodically to get network activity details.

### Aknowledgements

Special thanks go to the author of the
[Go package service](https://godoc.org/github.com/kardianos/service) without
whose work this implementation would be a real toil.

## Project Status

Alpha, first attempts to get things going.

## Project File Layout

```
root                    Go source files
|
+- pinstall             Installation scripts, helpers and notes
   |
   + - launchd          OS X
   |
   + - systemd          Linux
   |
   + - windows          Microsoft
```
