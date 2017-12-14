# dptool
dptool is a tool for working with IBM DataPower gateway appliances

Currently it supports just 3 primary commands:
+ write: Write file to gateway
+ delete: Delete file on the gateway
+ sync: Continuously syncs files from local dir to gateway dir (non-recursive, ignoring .* files)

The sync command is most useful during DataPower development of e.g. GatewayScript or XSLT allowing you to use Your Favorite IDE (complete with syntax highlighting and other goodies). All the while dptool continuously monitors the local filesystem for changes and synchronises them to a remote DataPower filesystem. Much better than editing files via the web-GUI or copying/pasting from your IDE to the web-GUI. Just hit save in your IDE and the changed are pushed.

## Installing

First of all download a binary suitable for your platform: [releases](releases),
or build from source as described later.

Either way, add the binary to your PATH.

## Running

dptool tries to help you, so just run without any arguments for clues:

    $ dptool
    For help, try: dptool help

Now ask for help:

    $ dptool help
    ... lots of nice help info...

Ask for help about the coolest command, sync:

    $ dptool help sync
    ... lots of nice help info about the sync command...

Examples:

    $ dptool write ~/tmp/hello.txt https://my_datapower_hostname:5554/mgmt/filestore/my_domain/local/hello.txt --user my_username --pass my_password

    $ dptool delete https://my_datapower_hostname:5554/mgmt/filestore/my_domain/local/hello.txt --user my_username --pass my_password

    $ dptool sync ~/tmp https://my_datapower_hostname:5554/mgmt/filestore/my_domain/local --user my_username --pass my_password
     
## Building From Source

This project currently depends on GB (https://getgb.io/) to manage dependencies and build.

So firstly, please install GB:

    $ go get github.com/constabulary/gb/...

Then it's a simple matter of running the following command in the root dir of this project:

    $ gb build

(a binary for your platform should have been compiled to the the bin/ directory)

Either add the bin/ dir to you PATH or just go:

    $ bin/dptool
    For help, try: dptool help

Now you know you have at least compiled a runnable binary, so go add it to you PATH to make life good again.

## Building For Other Platforms

If you want to build binaries for your friends on other platforms, it's just like using the standard go compiler, e.g.:

    $ GOOS=darwin GOARCH=amd64 gb build
    $ GOOS=linux GOARCH=amd64 gb build
    $ GOOS=linux GOARCH=386 gb build
    $ GOOS=windows GOARCH=amd64 gb build
    $ GOOS=windows GOARCH=386 gb build
    

(appropriately named binaries should now exist in the bin/ directory)

## Notes

1. Sorry, but AFIAK there is no way to make a GB project go-gettable :-(
