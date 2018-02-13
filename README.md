# dptool
dptool is a tool for working with IBM DataPower gateway appliances

Currently it supports just 3 primary commands:
+ write: Write file to gateway
+ delete: Delete file on the gateway
+ sync: Continuously syncs files from local dir to gateway dir (non-recursive, ignoring .* files)

The sync command is most useful during DataPower development of e.g. GatewayScript or XSLT allowing you to use Your Favorite IDE (complete with syntax highlighting and other goodies). All the while dptool continuously monitors the local filesystem for changes and synchronises them to a remote DataPower filesystem. Much better than editing files with the gui or copying/pasting from your IDE to the gui. Just hit save in your IDE and the changed are pushed.

## Installing

The quickest way to get up and running (if you have the Golang toolchain installed), is to build from source:

    $ go get github.com/samdjones/dptool

dptool should now be on your PATH (assuming $GOPATH/bin is on your PATH as usual).

If you can't build from source, download a binary suitable for your platform from [releases](../../releases). Your binary will be called something like 'dptool-linux-amd64' - it is suggested you rename it to simply 'dptool'. The rest of this doc assumes you did.

Either way, make sure the binary is on your PATH.

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
    PUT https://my_datapower_hostname/mgmt/filestore/my_domain/local/hello.txt 201 Created

    $ dptool delete https://my_datapower_hostname:5554/mgmt/filestore/my_domain/local/hello.txt --user my_username --pass my_password
    DELETE https://my_datapower_hostname:5554/mgmt/filestore/my_domain/local/hello.txt 200 OK

    $ dptool sync ~/tmp https://my_datapower_hostname:5554/mgmt/filestore/my_domain/local --user my_username --pass my_password
    PUT https://my_datapower_host:5554/mgmt/filestore/my_domain/local/hello.txt 201 Created
        ...hello.txt modified...
    PUT https://my_datapower_host:5554/mgmt/filestore/my_domain/local/hello.txt 200 OK
        ...hello.txt renamed to goodbye.txt...
    PUT https://my_datapower_host:5554/mgmt/filestore/my_domain/local/goodbye.txt 201 Created
    DELETE https://my_datapower_host:5554/mgmt/filestore/my_domain/local/hello.txt 200 OK
        ...goodbye.txt deleted...
    DELETE https://my_datapower_host:5554/mgmt/filestore/my_domain/local/goodbye.txt 200 OK

## Building For Other Platforms

If you want to build binaries for your friends on other platforms, it's really easy, e.g.:

    $ GOOS=darwin GOARCH=amd64 go build
    $ GOOS=linux GOARCH=amd64 go build
    $ GOOS=linux GOARCH=386 go build
    $ GOOS=windows GOARCH=amd64 go build
    $ GOOS=windows GOARCH=386 go build

## Notes

+ The sync feature is one-way - i.e. changes are pushed FROM local TO gateway only. This is not going to change anytime soon.
+ The sync feature will blindly overwrite whatever is on the gateway when first started - i.e. it does not look at last-modified times etc. **You have been warned**.
+ This tool isn't intended for management of production systems. Try [DPBuddy](https://myarch.com/dpbuddy) or [DCM](https://github.com/ibm-datapower/datapower-configuration-manager) instead. This tool is currently only aimed at helping DataPower *developement*, not *management*.
+ dptool has not been tested much on non-Mac platforms. There is a known issue on Windows (double-write when local file changes). I'll try and address this sometime soon. Come to think of it, I haven't tested on Linux for quite some time. **You have been warned**.
