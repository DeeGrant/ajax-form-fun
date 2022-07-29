# ajax-form-fun

Originally meant to test out some AJAX functionality, but became a test spot for other tools.
- [AJAX](https://www.youtube.com/watch?v=Wl-VvvNCk2A&list=PLWKjhJtqVAbkyK9woUZUtunToLtNGoQHB&index=10)
- [Mermaid](https://mermaid-js.github.io/mermaid/#/)
- Go serving static sites
- Go's [default ServeMux and Handlers](https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go)
- [jQuery](https://www.youtube.com/watch?v=KhtEmR2A1Fw&list=PLWKjhJtqVAbkyK9woUZUtunToLtNGoQHB)

### Run Project

In the terminal run `go run .` form the project directory.

Go to `localhost:8000` in your browser of choice.

### Learning Lessons

1. Go's default ServeMux doesn't require path prefixing sub-folder content like Gorilla's Mux [does](https://stackoverflow.com/questions/62499380/serving-static-files-in-an-http-server).

#### Mermaid Tests

[mermaid](https://mermaid-js.github.io/mermaid/#/)
: a very cool visual tool.

```mermaid
    graph TD;
        A-->B;
        A-->C;
        B-->D;
        C-->D;
```

```mermaid
    flowchart TD
    A[First step]-->B{Second Step};
    B -- Yes --> C[Option 1];
    B -- No --> D[Option 2];
    C --> E[Short throw];
    D ----> F[Long throw];
```

### Todo
- [Templates in static sites](https://www.alexedwards.net/blog/serving-static-sites-with-go)
- Marshall and Unmarshall with post?