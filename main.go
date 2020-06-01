package main

import (
	"fmt"
	"os"
	"flag"
	"path/filepath"
	"strings"
)

const (
	ZEROFILE_VERSION = "1.0"
)

const ( // error return values
	KErrorNone 		   int = 0
	KErrorFileNotFound int = 1
	KErrorTemp1		   int = 2
	KErrorTemp2	       int = 3
)

var ( // command line argument toggles
	muteFlagPtr 	   *bool // echo activity to console
	recursiveFlagPtr   *bool // recurse through folder structure?
	deleteFlagPtr	   *bool // delete matching files/folders?
	pruneFlagPtr       *bool // remove emoty folders?
	rootPath		   string // path to root search folder
)


func main () {
	getCommandLineArguments ()
	showBanner ()

	err := folderTreeScanner (rootPath)
	if err != KErrorNone {
		showError (err)
	}
}

func showBanner () {
	
	if !*muteFlagPtr {
		fmt.Printf ("ZeroFile Utility version %s\n\n", ZEROFILE_VERSION)
	}
}

func showError (err int) {
	var message string

	if !*muteFlagPtr {
		message += "Error : "
		switch (err) {
			case KErrorNone: // nothing
			case KErrorTemp1: 
							message += "Error Type 1"
			case KErrorTemp2:
							message += "Error Type 2"
			default:
				message += "Unknown error detected " + fmt.Sprintf ("(%d)",err)

		}
	}

	if err != KErrorNone {
		fmt.Println (message)
		os.Exit (err)
	}
}

func getCommandLineArguments () {
	
	muteFlagPtr = flag.Bool ("mute", false, "Echo activity to console")
	recursiveFlagPtr = flag.Bool ("recursive", true, "Recurse through folder structure")
	deleteFlagPtr = flag.Bool ("delete", false, "Delete all zero length files")
	pruneFlagPtr = flag.Bool ("prune", false, "Remove all empty folders and sub folders")
	flag.StringVar (&rootPath, "root", ".", "Path to top of search tree")

	flag.Parse ()
}


func folderTreeScanner (rootPath string) int {
	err := filepath.Walk(rootPath,
				 	     func(path string, info os.FileInfo, err error) error {
		  			         if err != nil {
		  			 	      return err
				   	        }
   	 		
					   	 // do whatever here
					   	 		
					   	 if info.Size () == 0 {
							// zero file or empty directory
							if *pruneFlagPtr {
								if info.IsDir () {
									
									if !*muteFlagPtr {
										fmt.Printf ("Deleting folder '%s' ...\n", strings.ToUpper(info.Name ()))
									}
									// delete folder and its sub folders
								}
							}

							// remove the file 
							if *deleteFlagPtr {
								if !*muteFlagPtr {
									fmt.Printf ("Deleting file : '%s'\n", strings.ToUpper (info.Name ()))
								}
								os.Remove (info.Name ())
							}
						}	
					   
					   	if err != nil {
					   		return err
					   	}

					   	return nil
				})

	if err != nil {
       	return KErrorTemp2
	}

	return KErrorNone
}




