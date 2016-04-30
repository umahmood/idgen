# IDGen

IDGen is a Go library which generates unique IDs.

# Installation

> go get github.com/umahmood/idgen <br/>

# Usage

See GoDoc for more information.

    package main

    import (
        "fmt"

        "github.com/umahmood/idgen"
    )

    func main() {
        fmt.Println(idgen.ID())

        id := idgen.New()

        for i := 0; i < 3; i++ {
            fmt.Println(id.ID())
        }
    }

Output:

    $ dad346dcb3a2e80ebe24ef2e4fd63a6486380427 
    $ 7d80fd18570d84f5b5d65629f0ddd5d2bc37a486
    $ acd2668ad4a444fcb8bc425b4e07fb5325ee044c
    $ cbf46eb3432d5c5ad1dc58c3f7836a424bac3982    

# Documentation

> http://godoc.org/github.com/umahmood/idgen

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
