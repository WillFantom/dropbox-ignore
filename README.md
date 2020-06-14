# Dropbox Ignore 

Git Ignore... But for Dropbox...

## Do I need the Dropbox CLI

Nope. This uses extended file attributes to flag files as ignored.

## How to use

Running the command `dropbox-ignore` will look for a `.dbignore` file in the current directory and ignores any file that matches the pattern (using `.` as the root). The pattern is **the same as for a git ignore file**.

If you already have a `.gitignore` file in a directory, simply run `dropbox-ignore -g` to us it.

## Install

### MacOS

Via Brew

```
brew tap willfantom/dropbox-ignore
brew install dropbox-ignore
```

## Is It Tested?

On MacOS: Yes... Otherwise No

### One Big Issue...

When a pattern is added to a gitignore, it will only impact matching files from when the change to the ignore file is made... This does not do that. When this application marks a file as ignored, it is also removed from your dropbox!