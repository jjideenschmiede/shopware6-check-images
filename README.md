# Shopware 6 image scraper

A small script that goes through all category pages and checks if all images are set on the following pages. At the end a file with all MPN and Urls is created.

If there are no missing images, then no file will be placed in the files folder.

## How to use?

The first thing you should do is clone the repository.

```console
git clone https://github.com/jjideenschmiede/shopware6-check-images.git
```

If the repository was cloned the config.json must be adapted. The URL should be specified. And the individual directory names must be stored in the categories array.

````json
{
    "url": "https://test.de/",
    "categories": [
        "Damenschuhe",
        "Hausschuhe",
        "Kinderschuhe",
        "Herrenschuhe"
    ]
}
````

If the previous points are all adjusted, then the program can be started. The best way to do this is to start a terminal and enter the following command.

```console
go run main.go
```

