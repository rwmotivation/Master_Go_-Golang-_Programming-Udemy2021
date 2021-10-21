Setup the Programming Environment on macOS (Go, Git and VSCode)
Step 1: Installing Go on macOS
a) Grab the latest version of Go for macOS from the official Go downloads page: https://golang.org/dl/

b) Open the package file you downloaded and follow the prompts to install Go.

c) Run the following command in the terminal to check the installed version of Go: go version

You should see an output similar to this: go version go1.17.1 darwin/amd64

Congratulations! You have installed Go on your system.



Step 2 — Creating Your Go Workspace
The Go workspace will contain three directories at its root:

pkg: The directory contains Go package objects compiled from Go source code, which are then used, at link time, to create the complete Go executable binary in the bin directory.

bin: The directory that contains executables built and installed by the Go tools.

src: The directory that contains Go source files. You’ll have a subdirectory of src for each Go application.

The default directory for the Go workspace as of 1.17 is your user’s home directory with a go subdirectory, or $HOME/go

Run the following command to create the directory structure for your Go workspace:

mkdir -p $HOME/go/{bin,src}

Note: Apple switched the default shell from bash to zsh in macOS Catalina.

Users updating older versions of macOS to Catalina receive a message asking them to switch to zsh as the default shell when they open the terminal.

If you are unsure which shell you use, run echo $0 or echo ${SHELL}



Edit ~/.bashrc or ~/.zshrc (if you are using zsh instead of the bash shell) file to set environment variables. Commonly you need to set 3 environment variables as GOROOT, GOPATH, and PATH.

Open ~/.bashrc or ~/.zshrc in your preferred editor and add:

export GOROOT=/usr/local/go

export GOPATH=$HOME/go

export PATH=$GOPATH/bin:$GOROOT/bin:$PATH



Now that you have the root of the workspace created and your $GOPATH environment variable set, you can create your future projects with the following directory structure.

mkdir $GOPATH/src/master_go_programming

Each Go program will reside in its own directory in $GOPATH/src/master_go_programming



Step 3 — Creating a Simple Program. Test the Installation.
a) Create a directory called hello_world in $GOPATH/src/master_go_programming

And inside that directory, a file called main.go



Write this sample program in main.go:

package main

import "fmt"

func main() {

fmt.Println("Hello Go World!")

}



b) Enable dependency tracking for your code.

Move to the directory with main.go and run: go mod init src/master_go_programming/hello_world



c) Run the test application

Move to the directory with main.go and run: go run main.go

You should see an output like this:

Hello Go World!



Installing Go on Mobile Devices

If you want to learn Go on your mobile phone just visit The Go Playground and write your code here: https://play.golang.org



Note: I'd strongly recommend going through this course using a laptop or a PC since we're going to work with a lot of code and keeping the pace is important.

Watching this training on a smartphone is not recommended because of readability issues.



Installing VSCode on macOS
Download and install the latest version of VSCode for macOS from https://code.visualstudio.com/Download

More details: https://code.visualstudio.com/docs/setup/mac



Setup VSCode for Go Programming
Open VSCode from Spotlight by typing visual studio.


Install Go Extension for VSCode

Go to File -> Preferences -> Extensions (Ctrl + Shift + x) and type go

Select go and install it.


3. After installing Go for VSCode click on Reload to Activate.

4. Important: Now exit completely from VSCode and open it up again. This step is required for it to install the extensions.

Note: Make sure you have Command Line Tools installed on your OS. If you are not sure launch the terminal (found in /Applications/Utilities/) and run: xcode-select --install

5. Go to View -> Command Palette (Command + Shift + P) and type goinstall

Then select ALL go tools and extensions and click ok.

The installation process can take a while and you know it’s ready when you see the message: All tools successfully installed. You are ready to Go :)

6. Optional: integrate the code command to the command line.

Open up the Command Palette by pressing Command + Shift + P and type shell. Select Install “code” command in PATH

7. Open your Go Workspace

Go to File -> Open Folder and select ~/go

In the src directory, there is a folder called master_go_programming that contains another folder for each Go program. Open main.go from hello_world folder you have just created



Whenever you see a popup, it's a chance to install something we may need.


Run the program by opening main.go in a terminal (In VSCode left-panel, right-click on main.go and Open in Terminal) and then type: go run main.go

Or you can run it by pressing Ctrl + F5.
