diggg
=====

Proceed a pool of goroutines which execute a `dig` command to a domain against
a list of given DNS servers addresses.

This is an experimental [go](http://golang.org/doc/install) project and it's an
wip as unit tests are missing.


Requirement
-----------

The `dig` command should be present and findable in `$PATH`.


Installation
------------

    $ go get github.com/toin0u/diggg
    $ go install github.com/toin0u/diggg


Configuration
-------------

The file `diggg.json` is the config file. It should be in the same folder like
the executable file `diggg`.


Usage
-----

    $ diggg google.dk


License
-------

diggg is released under the MIT License. See the bundled
[LICENSE](https://github.com/toin0u/diggg/blob/master/LICENSE) file for details.
