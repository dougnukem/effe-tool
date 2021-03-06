# Effe-tool, create and compile effes.

`effe-tool` is a simple command line utility that let you create and compiles new [effes][effe].

As you may ask what is an `effe`, an effe is an isolable computation unit, it is been build to be an open source alternative to AWS Lambda.

`effe-tool` gives you the possibility to manage fairly large projects completely written using `effes`.

Since `effe`s are state-less they provide you the foundation to build infinitely scalable systems, `effe-tool` provides the ability to manage those systems in a sane way.

## Getting Start

`effe` and `effe-tool` are built in go, I am assuming that you are not completely foreign to the languange.

You need to have `go` installed on your machine; if you type `go` in your terminal something should happen.

### Download effe-tool

Assuming your $GOPATH is set and that your $PATH contains $GOPATH/bin, then the quickest way to get `effe-tool` is: 

`go get github.com/siscia/effe-tool`

followed by 

`go install effe-tool`

Otherwise you can simply download the source file and compile it yourself.

## Create your first effe

To create your first `effe` all you need to do is `effe-tool new foo.go`. 

This command will created the file `foo.go`. 

Such file is already a valid `effe`, it serves as introductory example, but it is very scarse and you can easily bend it to your will.

``` go
simo@simo:~/gopath$ effe-tool new foo.go
Successfully created the new effe, path: foo.go
simo@simo:~/gopath$ cat foo.go 

package logic

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
		// it is possible to import any library in your gopath
)

var Info string = `
{
	"name": "hello_effe",  // these info will be used to 
	"version": "0.1",      // create the name of the executable
	"doc" : "Getting start with effe"
}
`

type Context struct {
	value int64
}

func Init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Start() Context {
	fmt.Println("Start new Context") 
	return Context{1 + rand.Int63n(2)}
}

func Run(ctx Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Hello from Effe:  %d\n", ctx.value)
	return nil
}

func Stop(ctx Context) { return }

```


## Compile your effe

Compile your `effe` is very simple as well. Continuing the example above all you need to do is `effe-tool compile foo.go`.

This simple command will compile your `effe`, the executable will be called as specified in the `Info` variable inside the source file.

It is also possible to specify in what directory put the executable, `--dirout dir_name` (default to `out/`) and how to call the executable, `--out name`, by default it will be called using the variable `Info` inside the source file.

``` bash
simo@simo:~/gopath$ effe-tool compile foo.go
File: foo.go | Everything went good, the file is been compiled.
Executable path: /home/simo/gopath/out/hello_effe_v0.1
simo@simo:~/gopath$ tree out/
out/
└── hello_effe_v0.1

0 directories, 1 file
```

## Compile a whole directory

It is also possible to compile a whole directory of `effe`s.

All you need to do is `effe-tool compile directory/` and `effe-tool` will try to compile every single file in that directory, it is pretty resistent to error and it will print all the information you may need.

When compiling a directory you can still provide the `--dirout dir_name` flag to decide in which directory put the resulting executable.

Also, keep in mind that compile preserve the folder structure of the source directory into the binary directory.

## Run your effe

Once your `effe` is been compiled you can run it, following the example above it is sufficient to run: `./out/hello_effe_v0,1`

Now if you go to `localhost:8080` you should see a welcome message.

``` bash
simo@simo:~$ curl http://localhost:8080
Hello from Effe:  2
```

## Contributing

Please.

You can simply open an issues or a pull request, I will do my best to reply promptly.

If you want to contribute but you don't know what to do just write me, I have more ideas than time.

## License

`effe-tool` is released under the MIT License, the same of `effe`.

[effe]: https://github.com/siscia/effe
