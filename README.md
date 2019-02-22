# ![alt text][logo]

dwm2go is a feature-rich statusline for the suckless [dwm](https://dwm.suckless.org/)
 window
manager. dwm2go is written in the Go Programming Language.

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

## Troubleshooting

**I do not see the symbols of the statusbar (Bitcoin symbol etc.)?**

*A: You have to install font-awesome on your system. The statusbar uses [font-awesome4](https://fontawesome.com/v4.7.0/cheatsheet/). Also add fontawesome to your config of dwm. To have the same font as in the screenshots you should have `static const char *fonts[]          = {  "White Rabbit:pixelsize=10;1", "fontawesome-webfont:pixelsize=8;1" };` in your config.h of your dwm installation.*

## Fonts
To get the same look please install the following fonts on your system and configure dwm accordingly.

* [White Rabbit](https://www.dafont.com/white-rabbit.font)
* [font-awesome4](https://fontawesome.com/v4.7.0/cheatsheet/)
