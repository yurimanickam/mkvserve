### Dependencies
---
 1. go - [site](go.dev)
 2. ffmpeg - [site](ffmpeg.org)
 3. mkvtoolnix - [site](mkvtoolnix.download)
##### Installation Commands
>[!WARNING]
>Do not arbitrarily run download links from random repositories. These links are here mostly for my own ease of use. I highly recommend you install these dependencies independently.

| OS                    | Command                                                                                                           |
| --------------------- | ----------------------------------------------------------------------------------------------------------------- |
| Linux (Ubuntu/Debian) | `sudo apt-get install -y golang ffmpeg mkvtoolnix`                                                                |
| Linux (Fedora)        | `sudo dnf install -y golang ffmpeg mkvtoolnix`                                                                    |
| Linux (Arch)          | `sudo pacman -S go ffmpeg mkvtoolnix-cli`                                                                         |
| Mac (Homebrew)        | `brew install go ffmpeg mkvtoolnix`                                                                               |
| Windows (winget)      | `winget install --id=Go.GoLang && winget install --id=Gyan.FFMPEG && winget install --id=MoritzBunkus.MKVToolNix` |
