# dptool
dptool is a tool for working with IBM DataPower gateway appliances

## Building

This project currently depends on GB (https://getgb.io/) to manage dependencies and build.

So firstly, please install GB:

    $ go get github.com/constabulary/gb/...

Then it's a simple matter of running the following command in the root dir of this project:

    $ gb build

(a binary for your platform should have been compiled to the the bin/ directory)

If you want to build binaries for your friends on other platforms, it's just like using the standard go compiler, e.g.:

    $ GOOS=windows GOARCH=amd64 gb build
    $ GOOS=windows GOARCH=386 gb build
    $ GOOS=darwin GOARCH=amd64 gb build

(appropriately named binaries should now exist in the bin/ directory)

Note: Sorry, but AFIAK there is no way to make a GB project go-gettable :-(
