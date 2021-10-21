Setup the Programming Environment on Linux (Go and VSCode)
Step 1: Installing Go on Linux (root permissions required)
Grab the latest version from the official Go downloads page: https://golang.org/dl/

On the website, you can find the URL for the latest binary release’s tarball, along with its SHA256 hash.



Open a terminal and move to your home directory or a directory with write access:

cd ~

Download the latest version of Go:

curl -O https://golang.org/dl/go1.17.1.linux-amd64.tar.gz

Next, extract the downloaded archive. It’s considered best practice to keep it under /usr/local:

sudo tar -xzvf go1.13.3.linux-amd64.tar.gz -C /usr/local

You will now have a directory called go in the /usr/local directory. Next, recursively change this directory’s owner and group to root:

sudo chown -R root:root /usr/local/go

Congratulations! You have installed Go on your system.



Alternatively, on Ubuntu 16.04+ you can install Go automatically from a repository.

In a terminal run the following commands:

sudo apt-get update

sudo apt-get install golang-go



If you are using a version of Ubuntu later than 16.04 and want to install the latest Go release you can use the longsleep/golang-backports PPA.



In a terminal run the following commands:

sudo add-apt-repository ppa:longsleep/golang-backports

sudo apt-get update

sudo apt-get install golang-go



Step 2 — Creating Your Go Workspace
The Go workspace will contain three directories at its root:



pkg: The directory contains Go package objects compiled from Go source code, which are then used, at link time, to create the complete Go executable binary in the bin directory.

bin: The directory that contains executables built and installed by the Go tools.

src: The directory that contains Go source files. You’ll have a subdirectory of src for each Go application.



The default directory for the Go workspace as of 1.17 is your user’s home directory with a go subdirectory, or $HOME/go, where $HOME is a variable that stores your home directory such as /home/john.



Run the following command to create the directory structure for your Go workspace:

mkdir -p $HOME/go/{bin,src}



Set $PATH and $GOPATH by adding the following lines to ~/.profile



Open ~/.profile in your preferred editor and add:



export GOPATH=$HOME/go

export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin



To update your shell, run the following command to load the global variables:

source ~/.profile



Now, check the installed Go version.

In a terminal run: go version



And you should receive an output similar to this one: go version go1.17.1 linux/amd64



Now that you have the root of the workspace created and your $GOPATH environment variable set, you can create your future projects with the following directory structure.



mkdir $GOPATH/src/master_go_programming



Each Go program will reside in its own directory in $GOPATH/src/master_go_programming



Step 3 — Creating a Simple Program. Test the Installation.
Create a directory called hello_world in $GOPATH/src/master_go_programming

And inside that directory a file called main.go

Write your sample program in main.go:



package main

import "fmt"

func main() {

fmt.Println("Hello Go World!")

}



Move to the same directory with main.go and run: go run main.go

You should see an output like this: Hello Go World!



Installing Go on Mobile Devices

If you want to learn Go on your mobile phone just visit The Go Playground and write your code there: https://play.golang.org



Note: I'd strongly recommend going through this course using a laptop or a PC, since we're going to work with a lot of code and keeping the pace is important.

Watching this training on a smartphone is not recommended because of readability issues.

Installing VSCode on Ubuntu Based Distributions
Download and install the .deb package (64-bit) from https://code.visualstudio.com/Download either through the graphical software center if it's available, or through the command line with:

sudo apt install 



# If you're on an older Linux distribution, you will need to run this instead:

# sudo dpkg -i

# sudo apt-get install -f # Install dependencies



Installing the .deb package will automatically install the apt repository and signing key to enable auto-updating using the system's package manager.



Then update the package cache and install the package using:

sudo apt-get install apt-transport-https

sudo apt-get update

sudo apt-get install code # or code-insiders

More details: https://code.visualstudio.com/docs/setup/linux



Setup VSCode for Go Programming
Open VSCode from the Menu or by typing code in a terminal.


Install Go Extension for VSCode

Go to File -> Preferences -> Extensions (Ctrl + Shift + x) and type go

Select go and install it.


Important: Now exit completely from VSCode and open it up again. This step is required for it to install the extensions successfully.



3. Go to View -> Command Palette (Command + Shift + P or F1) and type goinstall

Then select ALL go tools and extensions and click ok.



The installation process can take a while and you know it’s ready when you see the message: All tools successfully installed. You are ready to Go :)



4. Select the default shell in VSCode.

Go to View -> Command Palette (Command + Shift + P or F1) and type Select Default Profile.

Select bash as the default shell.



5. Open your Go Workspace

Go to File -> Open Folder and select ~/go

In the src directory there is a folder called master_go_programming that will contain another folder for each Go program. Open main.go from hello_world folder you have just created

Whenever you see a popup, it's a chance to install something we may need.


Run the program by opening main.go in a terminal (right-click on main.go and Open in Integrated Terminal) and then type: go run main.go

Or you can run it by pressing Ctrl + F5.
