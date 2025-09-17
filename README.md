# opds-list
List the files in the current directory in the [OPDS](https://opds.io/)
[format](https://en.wikipedia.org/wiki/Open_Publication_Distribution_System).
Currently only supports the title and url.

# usage

```
# build
go build

# install
cp opds-list /bin/

# use
opds-list [_site_ _prefix_]
site:   The domain of the site
Prefix: The url prefix for the files on the website.
```

For eg, if the file moby-dick.pdf is available on
https://mysite.tld/books/moby-dick.pdf, use:

opds-list https://mysite.tld /books/
