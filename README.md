# Chester

A minimal implementation of the Puppet forge API.

## Purpose
Take a directory full of Puppet module tarballs, serve it with nginx (or whatever), reverse proxy /v3/releases to Chester, and `puppet module install` to your hearts content. Chester is intended to implement the API endpoint(s) required to install a module using the Puppet module tool, and defer to another endpoint for retrieving the tarball.

## Installation
### Linux
```shell
sudo curl -L "https://github.com/jolshevski/chester/releases/download/1.0.0/darwin_amd64" -o /usr/bin/chester
sudo chmod +x /usr/bin/chester
```

### Mac
```shell
sudo curl -L "https://github.com/jolshevski/chester/releases/download/1.0.0/linux_amd64" -o /usr/bin/chester
sudo chmod +x /usr/bin/chester
```

## Usage
```shell
$ chester -help
Usage of ./chester:
  -binding string
        Golang ListenAndServe binding (default ":8080")
  -fileurl string
        URL to the base of the URL which the module tarballs are being served from (default "/v3/files")
  -modulepath string
        Directory containing module release tarballs to serve. Required.
```

## Contributing
If you find Chester useful, feel free to contribute enhancements / bug fixes.

### Testing
Unit tests are run with `make test`. The acceptance suite lives in the `acceptance` branch to simplify building with Travis.

### Releasing
Tagged commits will be built and uploaded to the associated release. The README's `Installation` section should be updated to reflect the new binary URL prior to release.
