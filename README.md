# latlong-cli
a simple CLI executable of [bradfitz/latlong](https://github.com/bradfitz/latlong) and its pre-built distribution of major platforms
- standalone executable
- lightweight: it's about 2.4MB after build
- minimal overhead: It takes just few milliseconds as like as `bradfitz/latlong`

# usage
```bash
$ ./latlong-cli 37.7833 -122.4167
$ America/Los_Angeles
```
## input
cli arguments (latitude, longitude)
## output
It's stdout. Keep that in mind there is no line break at the end of the result.

# pre-built distributions
- `linux/amd64` (linux x64)
- `darwin/amd64` (macOS x64)
- `windows/amd64` (Windows x64)

# more things
[node-latlong](https://github.com/Ecube-Labs/node-latlong): Node.js module wrapping `latlong-cli`
