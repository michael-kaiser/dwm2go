# ![alt text][logo]

dwm2go is a feature-rich statusline for the suckless [dwm](https://dwm.suckless.org/)
dwm window
manager. dwm2goline is written in the Go Programming Language.

## Screenshots
![alt text][full]

*dwm window manager with dwm2go statusline*

![alt text][small]

*the dwm2go statusline*

## Features

Up to now, dwm2go supports the following features:
* show date and time
* show battery status
* show number of updates for your system (uses `checkupdates`)
* show CPU utilization in percent
* show free space on your system in percent
* show status of mdp musicplayer and the current song (artist + title)
* show prices of cryptocurrencies (BTC, ETH)

[logo]: https://github.com/michael-kaiser/blob/blob/master/logo.svg "logo"
[full]: https://github.com/michael-kaiser/blob/blob/master/screenshot.png "fullscree"
[small]: https://github.com/michael-kaiser/blob/blob/master/justbar.png "just the toolbar"

## Dependencies
You need golang buildtools to build the files

For Arch install go with `sudo pacman -S go`

For Debian/Ubuntu do `sudo apt-get install go`

## Installation

1. Download dwm2go from github with `git clone https://github.com/michael-kaiser/dwm2go.git`
2. cd into the directory with `cd /directory/where/dwm2go/is/saved`
3. Build it with `go build`
4. run the binary with `./dwm2go`
5. to start at bootup, also add dwm2go to your *.xinitrc* file
