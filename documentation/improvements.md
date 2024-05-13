After 3 days (around 13.5 hours spent) and a functional project,
it is time to start thinking about some improvements and apply
them while we have time.

# Improve performance of the parsers

With current implementation, we loop over the html tokens 5 times
which takes a lot of time. On average, I can observe that the request
takes between 800ms and 1.2s. So ideally, we can improve this by doing 
the checks in one single loop. I don't think this would be doable
in the allocated time because I will have to change a big part of the logic
and it will require a lot of refactoring.
Another improvement would be to run parsers concurrently using goroutines
as following:

```go
// Create a wait group to wait for all goroutines to finish
var wg sync.WaitGroup
wg.Add(5) // Number of goroutines

// Run the function calls concurrently using goroutines
go func() {
    defer wg.Done()
    titleParser.Parse(string(body), analysis)
}()

go func() {
    defer wg.Done()
    htmlVersionParser.Parse(string(body), analysis)
}()

go func() {
    defer wg.Done()
    loginFormParser.Parse(string(body), analysis)
}()

go func() {
    defer wg.Done()
    headingsParser.Parse(string(body), analysis)
}()

go func() {
    defer wg.Done()
    linksParser.Parse(string(body), host, analysis)
}()

// Wait for all goroutines to finish
wg.Wait()
```

# Deployment

The most affordable way to host our solution would probably be hosting
the frontend with GitHub Pages or Cloudflare Pages, and the most affordable 
way to host a backend would probably be with Oracle's free Ampere VM or any 
other free tier part of solution like netlify or heroku. This will consume
some time so I will postpone it for now.

# Extensibility

We can look to extend our little project to extract more information like
number of `divs` or input labels. But with our implementation, it will be hard
to add new parser. We can utilize the factory pattern to solve this like 
following:

```go
type Parser interface {
	Parse(body string, analysis *Analysis)
}

// Factory function for creating parsers
func NewParser(parserType string) Parser {
	switch parserType {
	case "title":
		return &TitleParser{}
	case "htmlVersion":
		return &HTMLVersionParser{}
	case "loginForm":
		return &LoginFormParser{}
	case "headings":
		return &HeadingsParser{}
	case "links":
		return &LinksParser{}
	default:
		return nil
	}
}

...

// Create a wait group to wait for all goroutines to finish
var wg sync.WaitGroup

// List of parser types
parserTypes := []string{"title", "htmlVersion", "loginForm", "headings", "links"}

// Run the parser functions concurrently using goroutines
for _, parserType := range parserTypes {
    wg.Add(1)
    go func(pt string) {
        defer wg.Done()
        parser := NewParser(pt)
        if parser != nil {
            parser.Parse(string(body), analysis)
        }
    }(parserType)
}

// Wait for all goroutines to finish
wg.Wait()
```

# Performance optimization

Performance optimization in Go typically involves identifying and addressing bottlenecks in the code. 
First of all we can start by some profiling:
## 
Using Go's built-in profiling tools (pprof) we can identify hotspots in our code.
Profile CPU and memory usage to pinpoint areas that need optimization.

## Algorithmic Optimization
A straight example would be searching for the title only inside the head
part of the html and this would increase the performance

## Caching
We can use some temporal caching to avoid having too short successive calls.

## External Dependencies: 
We can also evaluate the performance of external dependencies and libraries 
used in the project. 