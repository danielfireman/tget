# tget
Super simple binary that prints a command based on ssh, scp and wget to fetch files via tunneling. Can be handy when working remotely and need to download files that can only be acessed via a firewall-protected network.

For instance, assuming if you would like to fetch http://foo.bar/fireman/public/myarticle.pdf but you could only access foo.bar ssh'ing fireman@bez:22. You could do this by running:

```sh
$ tget http://foo.bar/fireman/public/myarticle.pdf fireman@bez:22
Command to execute: ssh -t -o StrictHostKeyChecking=no fireman@bez:22 wget http://foo.bar/fireman/public/myarticle.pdf -O /tmp/myarticle.pdf; scp fireman@bez:22:/tmp/myarticle.pdf /tmp/myarticle.pdf
```

and then executing the printed command:

```sh
$ ssh -t -o StrictHostKeyChecking=no fireman@bez:22 wget http://foo.bar/fireman/public/myarticle.pdf -O /tmp/myarticle.pdf; scp fireman@bez:22:/tmp/myarticle.pdf /tmp/myarticle.pdf
```

The file is going to placed at /tmp.

## What is actually happenning?

Lets assume the following call:

```sh
$ tget http://foo.bar/fireman/public/myarticle.pdf fireman@bez:22 bizz
Command to execute: ssh -t -o StrictHostKeyChecking=no fireman@bez:22 'ssh -t -o StrictHostKeyChecking=no bizz wget http://foo.bar/fireman/public/myarticle.pdf -O /tmp/myarticle.pdf; scp bizz:/tmp/myarticle.pdf /tmp/myarticle.pdf'; scp fireman@bez:22:/tmp/myarticle.pdf /tmp/myarticle.pdf
```

And to make things easier, lets parse the reverse hosts and their command executions:

* bizz
    * wget http://foo.bar/fireman/public/myarticle.pdf -O /tmp/myarticle.pdf

* fireman@bez:22
    * ssh to bizz and execution of the commands listed above
    * scp bizz:/tmp/myarticle.pdf /tmp/myarticle.pdf

* locally
    * ssh to fireman@bez:22 and execution of the commands listed above
    * scp fireman@bez:22:/tmp/myarticle.pdf /tmp/myarticle.pdf
