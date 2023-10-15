Thanks to orignal [brokiem Automatic Hoyolab Check-in](https://github.com/brokiem/auto-hoyolab-checkin)
## Automatic Hoyolab Check-in 2

With this lightweight software, you don't have to worry about missing your daily check-in on the Hoyolab website because
this software will automatically check in to the website every 4 hours (Your PC must be on and connected to internet).

## How to use

1. Download the exe first in the release section or build yourself
2. Add the program shortcut to the auto startup program
   **Windows**
   `C:\Users\<YourUser>\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\`
   * You can also use the `Win+R` shortcut and type `shell:startup` to open the startup folder
3. Run the program and done, it will automaticaly run when your pc is turned on!

## Download

https://github.com/WeeraW/auto-hoyolab-checkin/releases/tag/new

## Build
1. Install latest golang from [here](https://go.dev/dl/)
2. Clone this repo
```sh
git clone https://github.com/WeeraW/auto-hoyolab-checkin.git
```
3. CD to project folder
```sh
cd auto-hoyolab-checkin
```
4. Resolve dependencies
```sh
go mod tidy
```
4. Build the executable
  * **Option 1:** Build for your current OS
     ```sh
     go build -o ./bin/hoyolab_auto_checkin.exe -ldflags="-s -w"  ./main.go
     ```
  * **Option 2:** Build without console (Windows only)
     ```sh
     go build -o ./bin/hoyolab_auto_checkin.exe -ldflags="-s -w -H=windowsgui"  ./main.go
     ```
5. You will get the executables with name `hoyolab-auto-checkin.exe`