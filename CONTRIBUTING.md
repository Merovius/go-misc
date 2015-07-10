Thank you for contributing. These are some general guideluines for your
contributions and communications in this repository:

Issues
===

Thank you for reporting an issue. To make dealing with them as simple as
possible, please consider the following things:
* Prefix the issue-title with the affected package name and a colon (i.e. "pkg:
  does not frobnicate correctly").
* Choose an expressive and meaningfull title to describe the problem.

For bugs:
* Describe the bug roughly. Give observed and expected behavior. If applicable,
  attach relevant logs.
* Create a minimum working example for reproducing the bug as either
  a gist, a go playground share or in some other online form. This minimum
  working example should be gofmt'ed.

For feature-requests:
* Describe the usecase you intend to solve. Code-Examples are always preferred,
  if it is real-world code, that is true in particular.
* Mention API-changes that might be required for your feature. If it breaks the
  API, try to find another way to include the feature. If you can't, don't
  worry, submit your issue anyway and I will think it through together :)

Pull-Requests
===

First of all, thank you for contributing code :) I try to have relatively high
standards on readability and maintainability and am relatively opinionated
about certain things. This means, that your PR is likely to take a couple of
iterations to get accepted. I appreciate your patience with this :)

Please make sure, the following is true, when you submit your PR:
* Any commit, that only changes a particular package, is prefixed with the
  package name. Look at existing commits for the exact format.
* In the end, your PR should get submitted as a single commit. In particular,
  this means that your PR should only contain changes, that make sense as a
  commit. If you want to make more than one independent change, please submit
  multiple PRs.
* All .go code is gofmt, go vet and golint clean. If you think that a
  particular warning by go vet or golint is not usefull, please meantion this
  in the description of your pull request (not the commit-message).
* API breaks are off-limits, unless for important bug-fixes. In particular,
  this includes
	* Adding, removing or changing the type of a method to an interface
	* Removing or changing the type of an exported Identifier
	* Changing semantics of a method or function

Coding Style
===

General coding style is idiomatic go. Readability has priority. Some
guideluines that you might take into account when contributing code (though it
is not necessary to know them all. Where applicable, they will be pointed out
in code review):
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

Conduct
===

In general, I expect respectfull conduct. If the go code of conduct is already
published, when you read this, you should apply it to your communications in
regards to this Repository. Otherwise, you might use the [Django Code of
Conduct](https://www.djangoproject.com/conduct/) as a rough guideluine.
