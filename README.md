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

To get the plaintext for a key:

```sh
$ fish showkey 1234 # key id is 1234 
```

To add an ssh key to your list of keys:

```sh
$ fish addkey newkey -p ~/.ssh/somekey.pub # the key's name is newkey
```

To remove a remote ssh key:

```sh
$ fish rmkey 1234 # key id is 1234
```

To create a droplet:
This create a droplet with the name droppy, the image id is 3668014,
the size is 512MB the region if New York 2 and you have login remotely
via ssh keys 1234 or 888 ( these are the remote key ids ).

```sh
fish create droppy -i 3668014 -s "512MB" -r "nyc2" -k 1234 -k 888
```

To destroy a droplet:

```sh
fish destroy 1111 # droplet id is 1111
```

The commands `reboot`, `on`, `off`, and `info` are similar to `destroy` in usage.  

To snapshot a droplet:
NOTE: the droplet should be powered off, do this with the `off` command.

```sh
fish snapshot 1111 -n snapshotname
```

To rebuild a droplet with a new image id (this keeps the ip address):

```sh
fish rebuild 1111 -i 4444 # 4444 is the new image id
```

The command `restore` is similar to rebuild, an important distinction is
restore can only use previous images or snapshots that machine had.

To resize a droplet:

```sh
$ fish resize 1111 -s "1GB" # 1GB is the new size of the droplet
```


### Installation

If you have Go on your system you can install it directly.

Binaries for platforms can be downloaded from the [Releases](https://github.com/Niessy/fish/releases).

### TODO

* Possible Fuzzy Matching on names instead of having to use ids

