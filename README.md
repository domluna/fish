fisherman
----------------

CLI for interacting with DigitalOcean.

Query from a url, similarly shaped in most cases.
Get JSON response and output it.

Can I do this with less repitition with an interface{} ?

TODO before viable:

1) fuzzy matching: doing things with ids is not fun and kind of
annoying.

2) when done put api in a seperate directory

4) Rethink configuration situation

Some observations:

The query parameters don't matter for the api only reads the parameters it needs for that particular endpoint. Thus I could have a generic string with all the parameters, help reduce code.
