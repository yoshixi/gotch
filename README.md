# ToDo
- [ ] Archive all assets as well as html so that it can load the web page locally with all assets loading properly.
- [ ] Use a browser engine such as puppeteer to enable scraping of javascript rendering pages.
- [ ] Identifying exception patterns and test them.
- [ ] Fix the test how to create your data for. https://github.com/yoshixj/gotch/pull/3#discussion_r790359761

# Usage

- With Go
```bash
$ go build -o gotch

$ ./gotch https://www.google.com https://autify.com
```

- With Docker
```bash
$ docker build . -t gotch
$ docker run --rm -it -p 80:80 gotch


/go/src/github.com/yoshixj/gotch # ./gotch https://www.google.com https://autify.com
/go/src/github.com/yoshixj/gotch # ./gotch --metadata https://www.google.com
site: www.google.com
num_links: 17
images: 1
last_fetch: Sun, 23 Jan 2022 17:52:05 UTC
/go/src/github.com/yoshixj/gotch # ./gotch --metadata https://autify.com
site: autify.com
num_links: 57
images: 104
last_fetch: Sun, 23 Jan 2022 17:52:05 UTC
/go/src/github.com/yoshixj/gotch # ./gotch --metadata https://autify.com  https://www.google.com
site: autify.com
num_links: 57
images: 104
last_fetch: Sun, 23 Jan 2022 17:52:05 UTC
site: www.google.com
num_links: 17
images: 1
last_fetch: Sun, 23 Jan 2022 17:52:05 UTC
```
