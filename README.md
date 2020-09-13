# Dropbox Ignore 

Git Ignore... But for Dropbox...

## Do I need the Dropbox CLI

Nope. This uses extended file attributes to flag files as ignored.

## How to use

Running the command `dropbox-ignore` will look for a `.dbignore` file in the current directory and ignores any file that matches the pattern (using `.` as the root). The pattern is **the same as for a git ignore file**.

If you already have a `.gitignore` file in a directory, simply run `dropbox-ignore -g` to use it instead of a `.dbignore` file.

## Install

### Github Release

> These are built via the action [here](.github/workflows/release.yml)

TODO: Create script to automate this

1. Download the release for your OS and architecture
2. Move the binary to a location in your `$PATH`
3. Ensure it is executable

### Brew [macos]

> probably don't use this yet...

```
brew tap willfantom/dropbox-ignore
brew install dropbox-ignore
```

## GitIgnore Inequalities

Dropbox Ignore has a key difference from gitignore. Adding a file to a gitignore will prevent git from staging future changes to the file however, the file state from the prior commit will be kept. Dropbox Ignore will simply flag the file as ignored. This will remove it from your dropbox if is already there! **TL;DR only add files to the dropbox ignore file that you want no trace of in your cloud.**

> Dropbox does store deleted files for a time, thus a mistake here may be recoverable.

## Is It Tested?

On MacOS: Yes... Otherwise No...
