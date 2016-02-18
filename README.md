System Searcher
==============
Description
--------------
<b>System Searcher</b> is a [Go](https://golang.org/) based desktop application, through which you can search in your computer's filesystem for same text into the files's contents. Apllication GUI is made with [go-gtk](http://mattn.github.io/go-gtk/). You can work directly with the system tree(copy, cut, delete and ect. files/dirs). The application offer directly view and edit files.

Installation
--------------

```
go get https://github.com/ElitsaVenchova/System-Searcher
```

Usage
--------------

[See GoDoc](https://godoc.org/github.com/ElitsaVenchova/System-Searcher) for
documentation and examples.

Search options
--------------
Search happens after pressing a "Start" button below a search form. "Cancel" button clean all search options(set them with their default values) without clean current search result. The different option for search in filesystem are:

* Search text field (mandatory)
* Type of search text field
    * Text (default)
    * Regular Expression
* Set search's root field
    * Empty - full search into filesystem  (default)
    * Set concrete path (inlude dir or file with mime type)

* Set search depth
    * Set max search depth  (0- only in current dir)
    * "Full search" (default)

Search result set
--------------

The result set is show like a list of files with mime type. When one of results is selected you can see below result:

* statistic info:
    * Ablosute file path
    * How many time is found seached text/matched regular expression
    * First 5 result - foreach show line with bold matched text
* Options for work with file:
    * "View" button - open File view/edit window
    * "Open" button - open file with default program
    * "Rename" button - remane file/dir with entered name into dialog box and pressed "Ok" button. "Cancel" button don't change selected file name
    * "Delete" button - delete selected file from filesystem after confirmation in dialog box
    * "Copy" button - copy selected file file
    * "Cut" button - cut selected file

System Tree
--------------

Systemtree show all subdirs and their files with some root. The dirs can be collapse and expand. For change the root and refresh a system tree must be pressed refresh icon in right of root field.

* Root field
    * Empty - show full system tree(default)
    * Concrete path - show tree with root this path(if path is incorect or file with mime type - the tree is empty)

* Option (clicked right button) on selected file/dir
    * Edit - active if selected item is a file. Open file in a File view/edit window
    * Copy - copy selected dir/file
    * Paste - active if have copied file/dir. Paste file into selected dir or into selected file's dir
    * Delete - delete selected dir/file
    * Select as root dir - change tree's root with seleted dir/file
    * Rename - remane file/dir with entered name into dialog box and pressed "Ok" button. "Cancel" button don't change selected file/dir name
    * Search in dir - set shearch root field with this path
    * Copy file path - copy the absolute file path

File view/edit window
--------------
In this window you can see and edit your file. File encoding is UTF-8.
Option below the showed file:

* "Save" button - set file content with content in a edit window
* "Original file" button - show original file content without save any changes
* "Cancel" button - cancel the view/edit window

License
--------------
This is free software, licensed under [The MIT License (MIT)](https://github.com/ElitsaVenchova/Sistem-Searcher/blob/master/LICENSE)

Contributing
--------------
Contributions welcome! Please fork the repository and open a pull request with your changes.
