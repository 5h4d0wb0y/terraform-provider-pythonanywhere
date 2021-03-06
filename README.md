PythonAnywhere Terraform Provider
=================================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Maintainers
-----------

This provider plugin is maintained by 5h4d0wb0y.

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Using the provider
------------------

Set the environments:

	export PYTHONANYWHERE_API_TOKEN="YOUR_API_TOKEN"
	export PYTHONANYWHERE_USERNAME="YOUR_USERNAME"

or set then directly in terraform configuration:

```
# For example, restrict provider version in 0.1.x
provider "pythonanywhere" {
  version   = "~> 0.1"
  username  = "YOUR_API_TOKEN"
  api_token = "YOUR_USERNAME"
}
```

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/5h4d0wb0y/terraform-provider-pythonanywhere`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:5h4d0wb0y/terraform-provider-pythonanywhere
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/5h4d0wb0y/terraform-provider-pythonanywhere
$ make build
```

Using the provider
------------------

```
# Configure the PythonAnywhere Provider
provider "pythonanywhere" {
  version   = "~> 0.1"
  username  = "YOUR_API_TOKEN"
  api_token = "YOUR_USERNAME"
}

```

See the [PythonAnywhere Provider documentation](https://help.pythonanywhere.com/pages/) to get started using PythonAnywhere.

Developing the Provider
-----------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-pythonanywhere
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
