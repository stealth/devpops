DevPops
=======

<p align="center">
<img src="https://github.com/stealth/devpops/blob/master/274px-ISO_7010_W009.svg.png" />
</p>


*DevPops* continues my research on companion worms. It is a friendly companion without payload and without modifying
any source files. The spreading vectors are gits.

In memoriam of my former Professor at the department of Computer Science, who was a great researcher and one of the
few who allowed his students virus experiments at his research lab.

**Disclaimer**: All of this code is for research and educational purposes only. By using this code, you agree to these
license terms and terms of use: **You are not allowed to modify or pass the source along to others. The sole purpose
and allowed use is to read and study the code and making experiments on an isolated lab machine without any network
connections. After the experiment, the lab setup has to be destroyed.**

If you want to experiment with *DevPops*, make sure:

* setup a *dedicated* lab VM or user-account
* do not have remote origins in the gits that can be found on the lab
* search the entire lab setup for ".git" directories and check each of them
  to contain no remote origins or origins that could be accessed by others locally
* plug off network cable and shutdown wifi
* after the experiment, destroy the lab

In order to prevent accidental clones or builds by drunken users or scripts, this repo cannot be
built directly after clone. It only contains non-functional, non-building code. It has to be edited by hand
in order to make sure the user understands what he is doing and that they are responsible for their own actions.

Additionally, in the event that you ignored my warnings above, there is a kill-switch built in:

```
localhost: $ touch ~/.devpops
```

and *DevPops* stops working.

Now that the intentions are very clear and you have been made aware of it, we can dig into the beauty.

The Code
--------

As my paper from 12y ago about companion worms disappeared from the server, I added it again to this repo 1:1 as it was
released in 2009. Inside its reference section, it contains a list of research that has been done about
such topics long ago. I want to put your focus on a thesis from 1980, four years before Fred Cohen made
his famous experiments.

In my paper from 2009, the companion used a terminal injection vulnerability (CVE-2008-2383) for spreading.
Git and golang were not very widespread back then. Due to the DevOps practise, git and cross-platform development
everywhere, the issue of companions becomes even more interesting. Tight coupling of Devel, Test and Release,
Code-Share and Collaboration-Everywhere create exactly an environment for companions that no longer need
vulnerabilities to spread. The "process" is doing that.

*DevPops* is very straight forward. It does not contain any payload, no mutation or encryption of its own code,
no mutation of the commit messages etc, although it would be dead easy to add. It just walks `$HOME` and `$GOPATH`
and puts a file called `devpops.go` in directories where it found existing golang files. It would then `git add`,
`git commit` and `git push`, taking advantage of git's feature to walk the current directory up until it found
a possible repo it could use.
Note that none of these are vulnerabilities or shortcomings, it just happens that (by design), all the DevOps
features are aligned so nicely, that it creates a *DevPops* opportunity for us. *DevPops* is sort of a quine,
the kind of fascinating programs, that print the source of itself. As such, it can even spread across architectures
and platforms. I leave the rest of the fun analyzing the details and reading the source to you.

I have to admit, I was too lazy to make a new paper out of it and I did not even research whether there is
similar stuff already in existence on vxheaven or elsewhere, as the idea seems very obvious. If you happen to have
your own research on it or make virii experiments at research labs or universities, feel free to give me a hint.

Resources
---------

The biohazard sign was taken from Wikipedia and is under public domain.

