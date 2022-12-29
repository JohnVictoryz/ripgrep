# RG aka ripgrep

The command that is better than grep\
Cross Platform File tools and downloader(like wget)\
*Also if you have a better slogan pls let me know*


*Feel free to contribute*\
*Started originaly by me: JohnVictoryz aka tutel for the Soviet Linux project*

# Buliding or runing
To build you need "go"
### For linux
### Debian/Ubuntu
```
sudo apt install go
```
### Fedora/Redhat
```
sudo dnf install go
```
### Arch/Manjaro
```
sudo pacman -S go
```
### Windows
Go to the [download page](https://go.dev/dl/)\
Donwload the .msi file run it
### macos
Go to the [download page](https://go.dev/dl/)\
Download the the .pkg and install it
## BUILDING
First all of the instructions are pretty much the same so its not a problem\
Optional Software: git\
Needed Software: go\
### Install git
### Debian/Ubuntu
```
sudo apt install git
```
### Fedora/Red hat
```
sudo dnf install git
```
### Arch/Manjaro
```
sudo pacman -S git
```
### Windows
Kinda complicated Not Recommended\
Go to the [github release of git for windows](https://github.com/git-for-windows/git/releases/latest)\
download the .exe and install it you will see further why this is going to be complicated
### macos
You will need to download Xcode and this will install git with it
## Cloning
### If you have installed git:
### Linux - macos:
Copy url and open a terminal\
write to the terminal: ``` git clone URL ```
### Windows:
Here is the trickey part you see git for windows has an emulated MINGW terminal Wich wont mount your drives with C:/ D:/ etc. it will have them on weird mount points the recommended way is to just right click on the folder you want and choose open git bash here or sth like that and type ``` git clone URL ```
### Without git:
### All OS:
Click the button that says code and download it as a zip and unzip it
### Build:
Go to the dir you cloned or unziped the repo and cd or go to src and type in a terminal or command promt ``` go built main.go ```
### Run:
The same but type instead of ``` go build main.go ``` ``` go run main.go ```\

*Feel free to contribute*\
*Started originaly by me: JohnVictoryz aka tutel for the Soviet Linux project*
