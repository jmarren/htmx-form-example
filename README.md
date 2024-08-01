# htmx-form-example
An example of how to create a form with HTMX.

## Usage
simply clone the repo, cd into the htmx-form-example directory, and run "air"

```
```bash
air
```
```
```
```





## Issue of Full-Page reload vs. Partial HTML

This is a common issue when first using HTMX. HTMX is not like Nextjs or other meta-frameworks. Its leaves you to do more of the basic setup. 

## The Render Function
To solve this issue, I use a function that can be found in /internal/render/render.go

You should use this function (or similar) when you are changing routes and want to be able to access the same route again. If the interaction is meant to be temporary/ephemeral you can simply send the HTML you'd like displayed.


