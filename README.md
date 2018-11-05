# Example Projects

## Topics we'll cover

1. Project structure design
1. Automation tooling
1. Setup project and tools (demo)

## Project Structure Design

> "It's turtles all the way down" - Urban Legend

Modern day web services are all about connecting many nuanced, self-contained pieces together, and hiding them behind the abstraction of a URL or rpc function.  A reliable, feature-rich, full-stack service likely has many such layers of abstraction.

> "Being abstract is something profoundly different from being vague -- The purpose of abstraction is not to be vague, but to create a new semantic level in which one can be absolutely precise" - Dijkstra

As such, beginning with an opinionated structure (the right abstraction layers) is incredibly important.  The practical benefit is, that if you start with a good project structure, that cleanly separates layers of abstractions, the code will write itself.  However, the art of creating a good project structure is time-consuming.  This is a classic `trade-off`; time invested in honing this craft today improves development velocity tomorrow.

> "Every skyscraper I've seen has a more or less rectangular base that narrows to a point – so we probably shouldn't hope to build one that looks like an inverted pyramid." - me

Thankfully (!), there are several huge (!!!), very complex open source projects written in Go, which serve as an excellent grounding point.  These projects **tend** to have carefully-considered standards of organization, and are absolutely the SWE equivalent of skyscrapers.  As examples, check out some of the packages in this [list](https://github.com/avelino/awesome-go).  Take notes on what structures the big ones have common.  In particular, I use Kubernetes and etcd as references all the time.

> "Perception is in the eye of the beholder" - Urban Legend

That being said, when you and I look at these repos, we will see different (!!) structures.  That's because they're HUGE and we will focus on different parts.  That's OK.  You just need something good enough for you, that is easy to communicate to others.  It's OK to have different opinions.  What I'll share with you is good enough for me today – check back in a week and it will have iterated, to be sure.

## Project Automation

> "Make me do something tedious twice, shame on you.  Make me do something tedious thrice, shame on me for not automating it out of my workflow." - me

Automation is a huge part of modern development.  Without certain tools, managing large codebases would be incredibly difficult.  We're going to discuss builds and tests.

> "We don't build skyscrapers by blowing up 4 floors and adding 5." - me

We want to be sure that every piece of code we add integrates with the existing codebase seamlessly.  Think of it like a dynamic prgrammnig problem – the incremental step needs to be safe and well-defined.  This means that we should not merge code that causes tests or builds to fail, and you don't add code without adding tests.

The tool we'll use as the foundation for this methodology is [Travis-ci](http://travis-ci.org/).

> "You don't build walls before floors, and floors before load-bearing beams, and load-bearing beams before a solid foundation."

You want to make sure that your foundations are sturdy.  To build antifragile systems, you need thoroughly-tested components.  That way, when your walls fall over (e.g. your new feature doesn't work as expected), you don't need to waste any time investigating your foundation.  This is another classic example of a `trade-off`.  Time invested in good tests today saves time debugging tomorrow.

We'll be using [codecov.io](http://codecov.io/) to track our tests, and to keep us honest on the tested parts of our code.  Thankfully, the built-in testing package for Go is really excellent, and is sufficient for generating the data necessary for `codecove.io`.  This means we don't sprinkle anything fancy in our code, we can pretend it's not even there. :)

