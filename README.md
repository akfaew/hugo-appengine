# Hugo on AppEngine
This is a skeleton for hosting your [Hugo](https://gohugo.io/) blog on Google AppEngine.

## Hugo
Hugo is a static site generator for blogs.
You write your articles in Markdown and export them to html.

## AppEngine
[Google AppEngine](https://cloud.google.com/appengine) allows you to run a Golang webapp for free.

## Configuration
See [the guide](https://jasminek.net/blog/post/hugo-appengine/) for details, but in short:
- Put your Hugo sources in `hugo/`
- Create a new AppEngine project (free)
- Edit the `Makefile` and set `PROJECT` to the newly created project name.

## Resources
- [Hosting Hugo on Google AppEngine](https://jasminek.net/blog/post/hugo-appengine/). An accompanying guide for this repository.
- [How to start a blog](https://syften.com/blog/post/blog/). Make an informed decision about the URL structure, SEO, RSS etc.
