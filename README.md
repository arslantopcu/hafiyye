# Hafiyye

## Getting Started

İnternette sizin belirttiğiniz keywordler ile periyodik arama yaparak size ilgili sonuçları listeler [Go language](http://www.golang.org/).

### Paketler:

    Revel Framework
    MongoDB

    go get github.com/revel/revel
    go get gopkg.in/mgo.v2
    go get github.com/PuerkitoBio/goquery
    go get github.com/janekolszak/revmgo



### Start the web server:

    revel run hafiyye

   Run with <tt>--help</tt> for options.

### Go to http://localhost:9000/ and you'll see:

"It works"

### Description of Contents

The default directory structure of a generated Revel application:

    myapp               App root
      app               App sources
        controllers     App controllers
          init.go       Interceptor registration
        models          App domain models
        routes          Reverse routes (generated code)
        views           Templates
      tests             Test suites
      conf              Configuration files
        app.conf        Main configuration file
        routes          Routes definition
      messages          Message files
      public            Public assets
        css             CSS files
        js              Javascript files
        images          Image files


![Screenshot](public/demo/1.png)

### Follow the guidelines to start developing your application:

* The README file created within your application.
* The [Getting Started with Revel](http://revel.github.io/tutorial/index.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/samples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).


## Contributing
We encourage you to contribute to Revel! Please check out the [Contributing to Revel
guide](https://github.com/revel/revel/blob/master/CONTRIBUTING.md) for guidelines about how
to proceed. [Join us](https://groups.google.com/forum/#!forum/revel-framework)!
