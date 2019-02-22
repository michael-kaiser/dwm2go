package main

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
// #include <unistd.h>
// #include <stdio.h>
// #include <stdlib.h>
// #include <stdarg.h>
// #include <string.h>
// #include <strings.h>
// #include <sys/time.h>
// #include <time.h>
// #include <sys/types.h>
// #include <sys/wait.h>
// #include <sys/statvfs.h>
// #include <X11/Xlib.h>
// Window getDefaultRootWindow(Display *dpy){return DefaultRootWindow(dpy);}
import "C"

import (
	"time"
	"os/exec"
	"fmt"
	"bytes"
	"os"
	"syscall"
	"net"
	"strings"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
)





func check(e error) {
    if e != nil {
        panic(e)
    }
}


func GetCryptoPrice(symbol string) float64{

	type Crypto struct{
		Usd float64`json:"USD"`
	}

	url :=
		"https://min-api.cryptocompare.com/data/price?fsym=" +
	symbol + "&tsyms=USD"

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print(err)
		return 0.0
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Print(getErr)
		return 0.0
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Print(readErr)
		return 0.0
	}

	crypto := Crypto{}
	jsonErr := json.Unmarshal(body, &crypto)
	if jsonErr != nil {
		log.Print(jsonErr)
		return 0.0
	}

	return crypto.Usd

}

func GetFreeMemory(mnt string) string{
	var stat syscall.Statfs_t
	syscall.Statfs(mnt, &stat)
	var total float64 = float64(stat.Blocks * (uint64)(stat.Frsize))
	var used float64 = float64((stat.Blocks - stat.Bfree) * (uint64)(stat.Frsize))
	return fmt.Sprintf(" %.0f%%", (used/total*100))
}
func SetStatus(name string, dpy *C.Display){
        var defaultroot = (C.Window)(C.getDefaultRootWindow(dpy))
	cs := C.CString(name)
	C.XStoreName(dpy, defaultroot, cs)
	C.XSync(dpy, 0)
}

func GetTime() string{
	dt := time.Now()
	time := dt.Format(" 15:04:05")
	return time
}

func GetDate() string{
	dt := time.Now()
	date := dt.Format(" 01-02-2006 Mon")
	return date
}

func GetUpdates() string{
	sep := []byte{'\n'}
	count := 0
	out, err  := exec.Command("checkupdates").Output()
	if err != nil{
		log.Print(err)
	}
	log.Print(out)
	count += bytes.Count(out, sep)
	return fmt.Sprintf("⟳ %d", count)
}

func GetCPULoadPercentage() string{
	var a [4]float64
	var b [4]float64
	file, err := os.Open("/proc/stat")
	check(err)

	defer file.Close()

	fmt.Fscanf(file, "cpu %f %f %f %f", &a[0], &a[1], &a[2], &a[3])

	time.Sleep(1 * time.Second)

	_, err = file.Seek(0,0)
	check(err)



	fmt.Fscanf(file, "cpu %f %f %f %f", &b[0], &b[1], &b[2], &b[3])

	loadavg := (100*((b[0]+b[1]+b[2]) - (a[0]+a[1]+a[2]))) / ((b[0]+b[1]+b[2]+b[3]) - (a[0]+a[1]+a[2]+a[3]));




	file.Close()
	return fmt.Sprintf(" %.2f%%", loadavg)
}

func nowPlaying(addr string) (np string, err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		np = "Couldn't connect to mpd."
		return
	}
	defer conn.Close()
	reply := make([]byte, 512)
	conn.Read(reply) // The mpd OK has to be read before we can actually do things.

	message := "status\n"
	conn.Write([]byte(message))
	conn.Read(reply)
	r := string(reply)
	arr := strings.Split(string(r), "\n")
	if arr[8] != "state: play" { //arr[8] is the state according to the mpd documentation
		status := strings.SplitN(arr[8], ": ", 2)
		if len(status) == 1{
			np = fmt.Sprintf("mpd - [%s]", status[0])
		}else{
			np = fmt.Sprintf("mpd - [%s]", status[1])
			//status[1] should now be stopped or paused
		}
		return
	}

	message = "currentsong\n"
	conn.Write([]byte(message))
	conn.Read(reply)
	r = string(reply)
	arr = strings.Split(string(r), "\n")
	if len(arr) > 5 {
		var artist, title string
		for _, info := range arr {
			field := strings.SplitN(info, ":", 2)
			switch field[0] {
			case "Artist":
				artist = strings.TrimSpace(field[1])
			case "Title":
				title = strings.TrimSpace(field[1])
			default:
				//do nothing with the field
			}
		}
		np = " " + artist + " - " + title
		return
	} else { //This is a nonfatal error.
		np = "Playlist is empty."
		return
	}
}

func GetBatteryStatus(base string) string{
	file, err := os.Open(base + "/present")
	check(err)

	defer file.Close()

	file, err = os.Open(base + "/charge_full")
	var descap int
	fmt.Fscanf(file, "%d", &descap)

	file, err = os.Open(base + "/charge_now")
	var remcap int
	fmt.Fscanf(file, "%d", &remcap)

	file, err = os.Open(base + "/status")
	var status string
	fmt.Fscanf(file, "%s", &status)

	switch status {
	case "Discharging":
		status = "-"
	case "Charging":
		status = "+"
	default:
		status = "f"
	}

	if(remcap < 0 || descap < 0){
		return "invalid"
	}

	return fmt.Sprintf(" %.0f%%%s",(float64)(remcap / descap) * 100, status)
}

type Data struct{
	cpu string
	memory string
	updates string
	battery string
	date string
	time string
	music string
	cryptoline string
}

func (d Data) StatusLine() string{
	return fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s %s",
	d.cryptoline, d.music, d.cpu,
	d.updates, d.memory, d.battery, d.date, d.time)
}

func CryptoLine(prices map[string]float64, symbols []string, logos []string) string{
	var line string
	for i,item := range symbols{
		line = fmt.Sprintf("%s%s %.2f$ ", line, logos[i], prices[item])
	}

	return line
}

func main() {
	var dpy = C.XOpenDisplay(nil)
	dataState := Data{music: "", cpu: "", memory: "", updates:
	"loading updates", battery: "", date: "", time: "",
		cryptoline: "loading price data"}

	var statusline string
	var cryptodata = make(map[string]float64)

	if dpy == nil {
		panic("Can't open display")
	}


	go func (){
		for ;; {
			batteryPath := "/sys/class/power_supply/BAT1"
			dataState.battery =
				GetBatteryStatus(batteryPath)
			dataState.memory = GetFreeMemory("/")
			time.Sleep(time.Minute * 5)
		}

	}()

	go func (){
		for ;; {
			time.Sleep(time.Second * 20)
			var symbols = []string{"BTC"}
			var logos = []string{""}
			for _,symbol := range symbols{
				cryptodata[symbol] = GetCryptoPrice(symbol)
			}

			dataState.cryptoline = CryptoLine(cryptodata, symbols, logos)
			time.Sleep(time.Minute * 5)
		}

	}()

	go func (){
		for ;; {
			time.Sleep(time.Second * 20)
			dataState.updates = GetUpdates()
			time.Sleep(time.Minute * 15)
		}
	}()

	for ;; {
		dataState.cpu = GetCPULoadPercentage()
		dataState.date = GetDate()
		dataState.time = GetTime()
		dataState.music, _ = nowPlaying("localhost:6600")
		statusline = dataState.StatusLine()
		SetStatus(statusline, dpy)
		time.Sleep(time.Millisecond * 50)
	}

}
