# basicproxy

basicproxy is a very simple proxy server for accessing blocked websites.  It only requests the file specified by the URL.  Therefore, HTML pages show up with broken links, images, stylesheets, etc.

## Why

I'm on a 14-hour trip from Gaspé to Montreal and the bus WiFi is blocking Imgur.

## Building

basicproxy is built to be used on the Google App Engine.  It's built in Go using [Negroni](http://github.com/codegangsta/negroni).  See the [App Engine documentation](https://developers.google.com/appengine/) for information on setting up your own version.
