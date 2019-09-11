# truspyd - A multi-platform service for tracking TCP, RAW and UDP sockets

Track socket details in order to gain network activity insights on a
per process/program level and make these insights available to the analytics
modules of a running [flowproc](https://github.com/shuntingyard/flowproc)
instance. 

## Implementation Notes

This software is intended to run as daemon/service on supported platforms
(Linux, Windows and OS X).

Currently existing utilities like ``ss`` or ``netstat`` are called periodically
to get network activity details.

## Project Status

Alpha, first attempts to get things going.

## Project File Layout

root                    Go source files
|
+- pinstall             Installation scripts, helpers and notes
   |
   + - launchd          OS X
   |
   + - systemd          Linux
   |
   + - windows          Microsoft
