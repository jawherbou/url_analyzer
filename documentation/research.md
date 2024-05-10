
This document will serve as a place to write down my brainstorming outcome and research done to implement requirements. It will also be updated proactively as new findings emerge. Updates may include additional functionalities, optimizations, or changes based on evolving best practices

# Objective
Web page that analyses web pages and by analysis we mean:
- html version
- Page title
- Headings and their levels
- Internal and external links and their accessibility 
- If the page contains a login form

Our page should contain a form (an input) where we provide the url and a modal or a textarea where we display results

# Possible solution
The main focus of our work and the difficult part would be get the web page analysis. And for that, I had two possible ideas:
- Using a third party library or api to get these information 
- Using a golang module that would parse the web page contents

First option was a dead end; no open source API available that can provide what we need. So I had to think about the second option. A quick search led me to find that using the net/http module I would be able to get the contents of a webpage from a url. Then using regex and searching for html tags in the contents would be doable. 

# Tools and POC
So now it is decided; we will start the POC parsing the contents of the response that we get from the net/http getUrl method.
The poc should only build a small api with one endpoint that will get the url in the query params and return the title of the webpage. I decided also to work with fibergo and not any other big framework. 

## Fibergo vs gin
For our tiny scope, there is no big difference between the two. Nevertheless, I chose to work with fibergo as uses a flexible routing system similar to Express.js, where routes are defined using HTTP methods and URL patterns. It provides a clean and intuitive way to define routes and handle requests.

Also the poc will be kept in a separate branch so that I can have a playground for other functionalities.