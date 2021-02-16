# Design

## Decisions
* Always skip the `.git` directory if it exists.
* Always Skip the `.` file.
* Scan runs from current directory. Not configurable. Change working directory to change scan directory.
* `.fend.yaml` configuration file. Found in working directory. Location not configurable, Name not configurable.
