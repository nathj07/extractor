# Extractor
[![CircleCI](https://circleci.com/gh/nathj07/extractor/tree/master.svg?style=svg&circle-token=0c9fb37da87f3dd9f9758a5c6fb279b626f760db)](https://circleci.com/gh/nathj07/extractor/tree/master)

## NOTICE: 
This library and it's API are currently subject to minor change. 

----


A simple library, with accompanying CLI to extract files from a `.tar.gz` archive.

What is particularly useful about both the library and the CLI is that you can specify not only the archive location and the unpack location, but also the file extensions you wish to unpack.

For example, consider the following archive - `~/myarchive.tar.gz`:
```
file_1.html
file_2.html
file_3.jpg
file_4.css
file_5.js
file_6.png
```
If you call the CLI as follows:

` ./extractor --source ~/myarchive.tar.gz -- dest ~/Desktop --exts "css,js"`

The result in the directory would be:
```
~/Desktop/myarchive/file_4.css
~/Desktop/myarchive/file_5.js
```

This can be very useful when you only want certain files from the archive and the rest can be discarded without ever unpacking them.

In oder to use this in your own projects you simply need to import this package and call the `targz.Extract` function. The tests provide clear examples of how to use this; there is no need to instantiate any new objects or manipulate any news structs. The function signature is very simple so you could even add an interface definition in your code that this would implement in order to allow you to mock this should you need to.

## Contributing
In the future this library may be extended to cover other archive formats. If you wish to contribute an `Extract` fuction please do so as follows:
- Make a fork
- Create a new package for the extractor, e.g. rar and add the new `extractor.go`
- Follow the signature of the existing `targz.Extract` function as this allows either this library, or a caller of the library, to define an interface for mocking in their own tests
- submit a pull request

It should be that all changes are non-breaking