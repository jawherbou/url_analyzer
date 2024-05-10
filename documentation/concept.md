
In this document, we will go over the project structure and 
the analysis details.

# Analysis
When we tokenize the html files, we get a stream of html tokens. 
Each token has a specific type as follows:
```go
const (
	// ErrorToken means that an error occurred during tokenization.
	ErrorToken TokenType = iota
	// TextToken means a text node.
	TextToken
	// A StartTagToken looks like <a>.
	StartTagToken
	// An EndTagToken looks like </a>.
	EndTagToken
	// A SelfClosingTagToken tag looks like <br/>.
	SelfClosingTagToken
	// A CommentToken looks like <!--x-->.
	CommentToken
	// A DoctypeToken looks like <!DOCTYPE x>
	DoctypeToken
)
```
## Getting Title

Here is an example of a title:
```html
<title> Title Example</title>
```

For getting the title here, we can read from the stream until 
we get a `StartTagToken` with `title` content and then the next element
in the stream would be `TextToken` with the title itself

## Getting HTML Version

After some research here is the html versions that we can have:
- HTML 2.0: <!DOCTYPE HTML PUBLIC "-//IETF//DTD HTML 2.0//EN">
- HTML 3.2: <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
- HTML 4.01 Strict: <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN">
- HTML 4.01 Transitional: <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
- HTML 4.01 Frameset: <!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN">
- XHTML 1.0 Strict: <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
- XHTML 1.0 Transitional: <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/- xhtml1-transitional.dtd">
- XHTML 1.0 Frameset: <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">
- XHTML 1.1: <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
- HTML 5: <!DOCTYPE html>

To get the html version, it is obvious that we 
can go over the stream until we get a `DoctypeToken`.

## Getting Headings and their levels

Same as getting the title, we search for `StartTagToken` tokens
that have h[1-6] tags and then we get the text in the next token.

## Internal and external links and their accessibility
In this part we have three parts to extract:

- Getting links: we would have three main types of links as follows:
```html
<a href="">
<script src="">
<link href="">
```
- If the link contains the host of the parent link/page, then that is an
internal. Otherwise, it's external.
- Accessibility: I would assume here that we mean if the link is well
formatted and valid or not


## If the page contains a login form
Here is an example of a form tag:

```html
<form action="URL" method="GET/POST" enctype="multipart/form-data">
    <!-- form elements go here -->
</form>
```

The search here should be also simple. We can search for the first
`StartTagToken` token containing `form`.


# Structure
The main focus will be on building the api. We will have a client 
but it will be worked on if I have time. And since we will work with
fibergo, we can have 

```
url_analyzer/
│
├───server/
│   ├───controllers/
│   │   └───webpage_parser_controller.go
│   │
│   │───middlewares/
│   │   ├───cors.go
│   │   └───query_validator.go
│   │
│   │───routes/
│   │   ├───not_found_route.go
│   │   └───webpage_parser_route.go
│   │
│   │───services/
│   │   ├───parsers
│   │   └───
│   │
│   │───main.go
│
├───client/
│

```

