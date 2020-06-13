# Dropbox Ignore 

Git Ignore... But for Dropbox...

## Why?

When building a complex program such as an OS, there can be a sh*t tonne of files built that you don't really need synced to dropbox. If they are synced, this can cause dropbox to use vast amounts of cpu cycles syncing the build files (that are often only temporary). This program will set all files that git ignores to also be ignored by dropbox sync.

### One Big Issue...

Dropbox Ignore will not only impact future changes to the files/folders, but actually remove the from dropbox all together!