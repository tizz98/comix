# Contributing

Contributions are more than welcome! The most common thing to contribute towards are new comic sources.
To add a new comic source, the `ComicDownloader` interface must be implemented.

## Example

Let's say we have a new "Foo Comic" source we want to add. Here are the steps to do so.

### Step 1 - Add type

In `app/downloader.go` we'd make the following change:
```diff
diff --git a/app/downloader.go b/app/downloader.go
index 7a09aa9..98d275c 100644
--- a/app/downloader.go
+++ b/app/downloader.go
@@ -25,6 +25,7 @@ type DownloaderType string
 const (
 	DownloaderTypeUnknown DownloaderType = "unknown"
 	DownloaderTypeXkcd    DownloaderType = "xkcd"
+	DownloaderTypeFoo     DownloaderType = "foo"
 )
 
 type DownloaderContext struct {
```

### Step 2 - Implement `ComicDownloader` interface

Create a new file for your comic source, in our case we'll create `app/foo.go` with the following contents:
```go
package app

import "net/http"

type FooDownloader struct {
	// add any config needed here
}

func (d *FooDownloader) DownloadComic(time.Time) (*Comic, error) {
	// add logic for determining comic image and title here
	return &Comic{
		ImageData: someBytes,
		Title: someTitle,
	}, nil
}
```

### Step 3 - Add new type to `cmd/downloader.go`

The next step is to use the new downloader type when requested. We'll make the following change to `cmd/downloader.go`:

```diff
diff --git a/cmd/downloader.go b/cmd/downloader.go
index 8d7d57c..73b5a8d 100644
--- a/cmd/downloader.go
+++ b/cmd/downloader.go
@@ -20,6 +20,8 @@ var downloaderCmd = &cobra.Command{
 		switch downloadSrc {
 		case "xkcd":
 			downloaderType = app.DownloaderTypeXkcd
+		case "foo":
+			downloaderType = app.DownloaderTypeFoo
 		}
 
 		outputDir, _ := cmd.Flags().GetString("output-dir")
```

### Step 4 - Run it!

Build the program and run it with your new source:

```bash
make build-pi
./comix-pi downloader --source foo
```

If you run into any problems, try running with the `--verbose` flag.

### Step 5 - Create a pull request

Make sure all tests are passing and request a review from @tizz98.
