# Design

## Decisions
* Always skip the `.git` directory if it exists.
* Always Skip the `.` file.
* Scan runs from current directory. Not configurable. Change working directory to change scan directory.
* `.fend.yaml` configuration file. Found in working directory. Location not configurable, name not configurable.
* Fail 0 byte size files. Add to skip configuration by extension or specific file name if not wanted.
* Not trying to identify and avoid binary files or other special files. The Failed File Extension output feature can help identify file types you may want to skip and then just configure the extension skips for your use-case.
* Not trying to implement a *fix* feature. Just reading and informing.
