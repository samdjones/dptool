package dptool

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// DPTool is a collection of IBM DataPower management tools.
type DPTool struct {
	Verbose bool
	User    string
	Pass    string
}

// NewDPTool is the factory function for DPTools.
func NewDPTool(verbose bool, user string, pass string) *DPTool {
	return &DPTool{verbose, user, pass}
}

// DeleteFile deletes a file on the gateway
func (dpt *DPTool) DeleteFile(remoteFileURL string) error {
	return doRequest(dpt.Verbose, remoteFileURL, dpt.User, dpt.Pass, http.MethodDelete, nil)
}

// PutFile PUTs a single local file to the gateway
func (dpt *DPTool) PutFile(localFilePath string, remoteFileURL string) error {
	bytes, err := ioutil.ReadFile(localFilePath)
	if err != nil {
		return err
	}
	payload := fmt.Sprintf("{\"file\":{\"name\":\"%s\",\"content\":\"%s\"}}",
		filepath.Base(localFilePath), base64.StdEncoding.EncodeToString(bytes))
	return doRequest(dpt.Verbose, remoteFileURL, dpt.User, dpt.Pass, http.MethodPut, strings.NewReader(payload))
}

// PutDir PUTs an entire local dir to the gateway (non-recursive, ignoring .* files)
func (dpt *DPTool) PutDir(localDir string, remoteDir string) error {
	// clean up optional trailing slashes
	localDir = filepath.Clean(localDir)
	// TODO: same for remoteDir

	// TODO: fail fast if localDir isn't a dir / doesn't exist
	// TODO: fail fast if remoteDir isn't a dir / doesn't existr

	files, err := ioutil.ReadDir(localDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		fileName := f.Name()
		if !isIgnorable(fileName) && !f.IsDir() {
			dpt.PutFile(filepath.Join(localDir, fileName), remoteDir+"/"+fileName)
		}
	}
	return nil
}

// SyncDir continuously syncs files from local dir to gateway dir (non-recursive, ignoring .* files)
func (dpt *DPTool) SyncDir(localDir string, remoteDir string) error {
	// clean up optional trailing slashes
	localDir = filepath.Clean(localDir)
	// TODO: same for remoteDir

	// TODO: fail fast if localDir isn't a dir / doesn't exist
	// TODO: fail fast if remoteDir isn't a dir / doesn't existr

	if err := dpt.PutDir(localDir, remoteDir); err != nil {
		return err
	}

	watcher, err := fsnotify.NewWatcher()
	defer watcher.Close()
	if err != nil {
		return err
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// fmt.Println(event)
				switch {
				case isIgnorable(event.Name):
					// ignore ignorables
					break
				case event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create:
					// it's a new or modified: PUSH remote file
					if fi, err := os.Stat(event.Name); err != nil || fi.IsDir() {
						// but ignore events about dirs or files we can't even stat (usually becasue they just got deleted)
						break
					}
					err = dpt.PutFile(event.Name, remoteDir+"/"+filepath.Base(event.Name))
				case event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename:
					// it's a rename or delete: DELETE remote file
					// TODO: figure out how to stop trying to delete non-existing dirs
					err = dpt.DeleteFile(remoteDir + "/" + filepath.Base(event.Name))
				}
				if err != nil {
					fmt.Println("sync error: ", err)
				}
			case err := <-watcher.Errors:
				fmt.Println("watch error: ", err)
			}
		}
	}()

	err = watcher.Add(localDir)
	if err != nil {
		return err
	}
	<-done
	return nil
}

// figure out if file should be ignored during sync (just ignore .* files for now)
func isIgnorable(fileName string) bool {
	r, _ := regexp.Compile("^[.]")
	return r.MatchString(fileName)
}

// TODO: move to generic util package
func doRequest(verbose bool, url string, user string, pass string, method string, data io.Reader) error {
	fmt.Printf("%s %s ", method, url)

	// make InsecureSkipVerify a command option?
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return err
	}

	req.SetBasicAuth(user, pass)
	// TODO: maybe handle timeout a bit better?
	response, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println(response.Status)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	if err != nil || !(response.StatusCode == 200 || response.StatusCode == 201) {
		// print resp headers and body content in verbose mode to help debug issues
		if verbose {
			hdr := response.Header
			for key, value := range hdr {
				fmt.Println(key, ":", value)
			}
			fmt.Println(string(contents))
		}
		if err == nil {
			err = errors.New(response.Status)
		}
		return err
	}
	return nil
}
