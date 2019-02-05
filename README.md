# tyler-website
#### My personal website

## Purpose
This website was built for two reasons, mostly.
* I wanted a portfolio, of sorts.
* I wanted to learn Go in a more real-world situation.

## Requirements for deployment
* Go 1.11
* A means of connecting to a database (in my case, a super secret *.yaml file I have only on my Appengine server)

## Hosting
My website as of this Readme's creation (February 2, 2019) is hosted on Google Cloud's Appengine.  

## Notes
I have this habit of reinventing the wheel, so I ended up not using any frameworks
in favor of writing the whole thing from scratch; not counting a few libraries here and there to save me from some gray hairs.   

I wrote a barebones CMS (for lack of a better word) to save me from needing to redeploy every time I tweak a page,
and in theory, this works quite well. Practice, however, is another story:  

My text data is stored in and retrieved from a separate database, so I can easily manipulate that through my on-site editor.  

The other portion of my site manager is a file editor, which allows for tweaking of front-end files (.html, .css, .js, etc). 
Testing on localhost, this works flawlessly.  However, after finally deploying to my site, this portion would not work. 
As it turns out, I found that Appengine does not allow for modification of files in their deployed
state, rendering that weekend of work useless, unless I find another host.

This is what it looks like in action:
[![data manager](https://github.com/tston529/tyler-website/raw/master/references/edit.png)](#cms)
Shame saving files won't work on Google's Appengine platform, I was especially proud of the directory tree parser 
I wrote to make things look nice and clean on the left (see scripts/edit.js)
