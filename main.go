// License: AGPL-3.0-or-later
package main

import (
	"fmt"
	"html/template"
	"log"
	"mime"
	"os"
	"path/filepath"
)

type info struct {
	Title, Link, Mime string
}

type container struct {
	Site  string
	Infos []info
}

func getOpds(dir, prefix string) ([]info){
	d, err := os.ReadDir(dir)
	if err != nil {
		log.Panic(err)
	}
	i := []info{}
	for _, f := range d {
		name := f.Name()
		if f.IsDir(){
			o := getOpds(dir + "/" + name, prefix + name + "/")
			i = append(i, o...)
			continue
		}
		t := mime.TypeByExtension(filepath.Ext(name))
		ii := info{Title: name,
			Link: prefix + name,
			Mime: t,
		}
		i = append(i, ii)
	}
	return i
}

func usage(){
	u := `
opds-list [site prefix]
site:   The domain of the site
Prefix: The url prefix for the files on the website.

For eg, if the file moby-dick.pdf is available on
https://mysite.tld/books/moby-dick.pdf, use:
opds-list https://mysite.tld /books/
`
	fmt.Fprintln(os.Stderr, u)
}

func main(){
	site:="http://192.168.18.4"
	prefix:="/"
	if len(os.Args) > 2 {
		site=os.Args[1]
		prefix=os.Args[2]
	} else if len(os.Args) > 1 {
		usage()
		return
	}
	c := container{Site: site, Infos: getOpds(".", prefix)}
	t, err := template.New("opds").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	// html/template assume's you're and idiot and replates < and >
	// that appear in _template_, not the data. <?xml -> &lt;xml
	fmt.Fprintln(os.Stdout, `<?xml version="1.0" encoding="UTF-8"?>`)
	t.Execute(os.Stdout, c)
}

var tmpl = `
<feed xmlns="http://www.w3.org/2005/Atom"
      xmlns:dc="http://purl.org/dc/terms/"
      xmlns:opds="http://opds-spec.org/2010/catalog">

  <title>My Library</title>
  <updated>2010-01-10T10:01:11Z</updated>
  <author>
    <name>Me</name>
    <uri>{{.Site}}</uri>
  </author>
 {{range .Infos}}
  <entry>
    <title>{{.Title}}</title>
    <link rel="http://opds-spec.org/acquisition" 
          href="{{.Link}}"
          type="{{.Mime}}"/>
  </entry>
  {{end}}
</feed>
`
