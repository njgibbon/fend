# fend
Check for Newline at **F**ile **End**

Fend recursively scans all files in a directory and reports those that don't end in a new line character (`\n`).

TODO: Input GA Output picture.

# Why might you want to do this?

This StackOverflow post captures the *why?* better than I could:

https://stackoverflow.com/questions/729692/why-should-text-files-end-with-a-newline

It will mean no more GitHub warnings for 'No newline at EOF' on Pull Request.

TODO: Image

By enforcing this check using the GitHub Action you can automate a basic Standard in your project.

Consistent is clean. Clean is good. Don't overthink it. :ok_hand:

# GitHub Action
TODO

# Details
* Fend always ignores all '.git' directories. To skip anything else see **Configuration**.

# TODO
## Docs
* Readme docs.
* GH Action release.
* Research doc with examples of binary skip configs.
## Code
* Finish testing.
* Define and use data structures for ScanConfig and ScanResult.
* Time the scan and add this to ScanResult.
* Percentages to ScanResult and output.
* Target mode feature.

# Usage
```
go get github.com/njgibbon/fend
# ensure binary can be foud in $PATH
cd <dir-to-scan>
fend
```
**Other**
```
# Version of fend
fend version
# Doc command points to here
fend doc
```
## Package
TODO

# Configuration
TODO

# Failed Scan Example
TODO: Image

# Meta
Project used as a vehicle to help learn some of the basics of tool development using GoLang and also to explore GitHub Actions from a development point of view as I have had a really positive experience with GA as a User.

# Similar tools
TODO: write nice things about the other tool I found.
