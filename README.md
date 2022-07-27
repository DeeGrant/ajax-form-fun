# ajax-form-fun

Originally meant to test out some AJAX functionality, but became a test spot for other tools.
- AJAX
- [Mermaid](https://mermaid-js.github.io/mermaid/#/)
- Go serving static sites
- Go's [default ServeMux and Handlers](https://www.alexedwards.net/blog/an-introduction-to-handlers-and-servemuxes-in-go)


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
