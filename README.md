# htmx-form-example
An example of how to create a form with HTMX.

## Usage
simply clone the repo, cd into the htmx-form-example directory, and run "air"

```bash
air
```

The application is run on http://localhost:3000 by default

## Issue of Full-Page reload vs. Partial HTML

This is a common issue when first using HTMX. HTMX is not like Nextjs or other meta-frameworks. Its leaves you to do more of the basic setup, but allows for much greater flexibility.

## The Render Function
To solve this issue, I use a function that can be found in /internal/render/render.go

You should use this function (or similar) when you are changing routes and want to be able to access the same route again. If the interaction is meant to be temporary/ephemeral you can simply send the HTML you'd like displayed.

## How the Render function works

I have created a base.html file and a sign-in-form.html file in /ui/templates. If a request is made by HTMX, the Hx-Request Header will be included in the request. If this header is present, the Render function will send only the html template that is passed to it as the partialTemplate arg, rendered with the data that is passed it. However, if  full page reload occurs, the Hx-Request header will not be present. If this is the case, the Render function will pass the name of the partial template into the base.html template and send back the entire page with the partial template included.

## Things to Remember

If you want the new page to be pushed onto the history stack, it is important to do one of the following: 
  - set the hx-push-url attribute to true on the requesting element 
  (this will push the requested url onto the history stack)
  - set hx-push-url to another url that you would like pushed (i.e. hx-push-url="/sign-in-success")
  - If you are redirecting server-side, depending on some server-side logic, it is probably best to use the Hx-Push-Url header on your response so that the correct url is pushed onto the stack depending on your server-side logic.

Naming is important. If you find that some elements are being replaced into the same position over and over again. It is probably best to assign a div (or other element) to mark that position with one consistent id, then use it as the target for each route.

