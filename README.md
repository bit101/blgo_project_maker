deprecated. using https://github.com/bit101/tinfox

# blgo_project_maker

Simple utility to create a new blgo project.

Usage:

`./newproject <path>`

Example:

```
cd blgo_project_maker
./newproject ../myproject
cd ../myproject
make
```

Warning: there's not a lot of protection on the location of the new project. It's just going to try to run `mkdir` on the path you give it and then copy files into that dir. If there's something already at that path, it will probably just fail, but use at your own risk.
