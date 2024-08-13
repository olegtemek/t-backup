# t-backup
This app was created primarily for me. For local synchronization and for synchronization with github

<br>

## Getting started
1. Download application for your OS [Link](https://github.com/olegtemek/t-backup/releases/latest)

2. Available flags <br>

| Flag name        | Description                                               | Example                      |
|------------------|-----------------------------------------------------------|------------------------------|
| name             | backup name                                               | vault                        |
| numberOfBackups  | number of backups                                         | 3                            |
| originalPath     | Path to the folder you want to back up                    | /users/test/important_folder  |
| driver           | currently only 2 drivers: `git` and `local`                   | local                        |
| backupPath       | Path to where backups are saved, if you used the `local` driver | /Users/test/backups/          |

3. Then you can try to start
```
./t-backup --name=vault --numberOfBackups=3 --orinialPath=/Users/test/important_folder --driver=git
```

> Notes: can't remember if you are using the `git` driver, then you need to initialize git and run push origin

> At the moment only one operating system is supported - Mac OS, but you can also try it under others

<br>

## How to automate startup and add to your scheduler

### Macos

For Macos, you can use this sample .plist file [Link](https://github.com/olegtemek/t-backup/blob/main/examples/example.plist). This script will be start sync every day

### Windows

later

### Linux

later