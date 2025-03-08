![logo](https://github.com/user-attachments/assets/0bbc4877-7d9c-49f3-8253-15f1c87bd422)
<p align="center">
<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white">
<img src="https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white">
<img src="https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white">
<img src="https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB">
<img src="https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white">
<img src="https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black">
<img src="https://img.shields.io/badge/mac%20os-000000?style=for-the-badge&logo=macos&logoColor=F0F0F0">
</p>

### CloakCrypt allows you to encrypt and hide the contents of a file inside another file in a stealthy manner. 

Think of it as a more stealthy version of an encrypted archive. The encrypted container will act just as if it were the original file, but in reality, it contains the encrypted data. Even if opened in a hex editor, it is difficult to spot if it has encrypted data or not due to CloakCrypt adding the last segment of the original file to the end of the encrypted part so the ending looks natural.

## Installation
* Windows: download the exe from the releases tab
* Linux and MacOS / manual compile:
```
git clone https://github.com/SpoofIMEI/CloakCrypt
cd CloakCrypt

#Install golang if need be (linux: apt install golang)
go install github.com/wailsapp/wails/v2/cmd/wails@latest
~/go/bin/wails doctor #If everything looks fine continue to next step
npm install -g sass
sass frontend/src/
~/go/bin/wails build
# The executable can be found in build/bin
```

## Support
The program will probably work with other filetypes as well, but here are ones I personally tested it with:
* <b>ZIP</b>
* <b>RAR</b>
* <b>EXE</b>
* <b>TXT</b>

Supported platforms:
* <b>Windows</b>
* <b>Linux</b>
* <b>Mac OS</b>

## Screenshots:
### In app:
<img width=400 src="https://github.com/user-attachments/assets/f6f55a47-774f-420c-8819-8201ba7a3dd02"></img>
<img width=400 src="https://github.com/user-attachments/assets/bc5fae2c-650e-4be1-a6de-b68d6b63d462"></img>


### Encrypted container inside PSEXEC:
<img width=800 src="https://github.com/user-attachments/assets/4605f29a-18fa-41d4-a4d7-804687aa34be"></img>

## Technical specifications
```
GUI: Wails
Main language: Go

Encryption: AES-256 GCM
Key: Salted Shake256 sum of at least 7 character password
```
