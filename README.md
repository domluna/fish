fish
----------------
fish is a Command Line Tool for interacting with DigitalOcean.

In order to use fish you'll need your api key and client id. These
can be found [here](https://cloud.digitalocean.com/api_access).

Note that you can only view an api key ONCE, so make sure you write
it down somewhere.

### Configuration

fish reads your config info from $HOME/.fish. See sample.config for a
sample configuration.

Alternatively you can use environment variables DIGITALOCEAN_CLIENT_ID
and DIGITALOCEAN_API_KEY

### Examples

To view all your droplets:

```sh
$ fish droplets
$ ...
```

To view available sizes:

```sh
$ fish sizes
$ ...
```

To view available regions:

```sh
$ fish regions
$ ...
```

To view images:

```sh
$ fish images # these are your snapshots
$ ...
$ fish images -g # this enables all default DigitalOcean images
```

To view ssh keys:

```sh
$ fish keys
$ ...
```

### Installation

If you have Go on your system you can install it directly.

Binaries for platforms can be downloaded from the [Releases](https://github.com/Niessy/fish/releases).

### TODO

* More examples
* defaults
* Possible Fuzzy Matching on names instead of having to use ids

