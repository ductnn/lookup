# LOOKUP

[![CI](https://github.com/ductnn/lookup/actions/workflows/ci.yml/badge.svg)](https://github.com/ductnn/lookup/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ductnn/lookup)](https://goreportcard.com/report/github.com/ductnn/lookup)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/ductnn/lookup/pulls)
[![license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**LOOKUP** is a simple `reconnaissance` tool, a tool for scanning IP informations,
CNAME, Subdomain, NS(Name Server), A records, MX records, ... with **Go**.

```bash
               #   ___       #   ___           ___                    #                 #
     _/7       #  <_*_>      #  <_*_>         /\#/\         \-^-/     #=ooO=========Ooo=#
    (o o)      #  (o o)      #  (o o)        /(o o)\        (o o)     #  \\  (o o)  //  #
ooO--(_)--Ooo--8---(_)--Ooo--8---(_)--Ooo-ooO--(_)--Ooo-ooO--(_)--Ooo---------(_)--------
```

## Install

First, install [golang](https://go.dev/doc/install).

Then, clone from soucre code and setup:

```sh
git clone https://github.com/ductnn/lookup.git
cd lookup
go get
```

## Usage

### Without docker

Run command:

```
go run main.go
# Enter domain you want to check.
# Example
google.com
```

  <p align="center">
    <img src="assets/demo.png" width="500">
  </p>

Check result:

**CNAME**

  <p align="center">
    <img src="assets/cname.png" width="500">
  </p>

**TXT records**

  <p align="center">
    <img src="assets/txt.png" width="500">
  </p>

**IP informations**

  <p align="center">
    <img src="assets/ip.png" width="500">
  </p>

**Name Servers**

  <p align="center">
    <img src="assets/ns.png" width="500">
  </p>

**MX records**

  <p align="center">
    <img src="assets/mx.png" width="500">
  </p>

### With docker

Check [Dockerfile](Dockerfile) and build with command:

```sh
docker build -t <your-image> -f Dockerfile .
```

or pulls my image in [here](https://hub.docker.com/r/ductn4/loo), and run *container*

```sh
docker run -it ductn4/loo:latest
```

### With file binary

You can down file *binary* in [here](bin/loo) for running tool:

```sh
# Download file bin to your local
chmod +x loo
./loo
# And check result ...
```

Test with **github.com**:

```sh
➜  lookup git:(master) go run main.go

               #   ___       #   ___           ___                    #                 #
     _/7       #  <_*_>      #  <_*_>         /\#/\         \-^-/     #=ooO=========Ooo=#
    (o o)      #  (o o)      #  (o o)        /(o o)\        (o o)     #  \\  (o o)  //  #
ooO--(_)--Ooo--8---(_)--Ooo--8---(_)--Ooo-ooO--(_)--Ooo-ooO--(_)--Ooo---------(_)--------


Enter subdomain or domain name:
github.com

Enter your choice:

[1] - CNAME lookup
[2] - Subdomain lookup
[3] - TXT Records lookup
[4] - IP information lookup
[5] - NameServers lookup
[6] - MX Records lookup
[7] - Extract URL
[8] - Lookup all without subdomains
[9] - ALL

2

[✔] Subdomain
+------------+-----------------------------------+
|   DOMAIN   |             SUBDOMAIN             |
+------------+-----------------------------------+
| github.com | import2.github.com                |
+            +-----------------------------------+
|            | import2.github.com                |
|            | importer2.github.com              |
|            | porter2.github.com                |
+            +-----------------------------------+
|            | api.stars.github.com              |
+            +-----------------------------------+
|            | api.stars.github.com              |
|            | www.api.stars.github.com          |
+            +-----------------------------------+
|            | *.github.com                      |
+            +-----------------------------------+
|            | *.github.com                      |
|            | github.com                        |
+            +-----------------------------------+
|            | examregistration.github.com       |
+            +-----------------------------------+
|            | vpn-ca.iad.github.com             |
+            +-----------------------------------+
|            | *.review-lab.github.com           |
+            +-----------------------------------+
|            | *.review-lab.github.com           |
|            | review-lab.github.com             |
+            +-----------------------------------+
|            | github.com                        |
+            +-----------------------------------+
|            | github.com                        |
|            | www.github.com                    |
+            +-----------------------------------+
|            | *.pkg.github.com                  |
+            +-----------------------------------+
|            | *.pkg.github.com                  |
|            | pkg.github.com                    |
+            +-----------------------------------+
|            | visualstudio.github.com           |
+            +-----------------------------------+
|            | visualstudio.github.com           |
|            | www.visualstudio.github.com       |
+            +-----------------------------------+
|            | *.github.com                      |
|            | github.com                        |
|            | www.github.com                    |
+            +-----------------------------------+
|            | render-lab.github.com             |
+            +-----------------------------------+
|            | render-lab.github.com             |
|            | www.render-lab.github.com         |
+            +-----------------------------------+
|            | *.github.io                       |
+            +-----------------------------------+
|            | docs-front-door.github.com        |
+            +-----------------------------------+
|            | docs.github.com                   |
+            +-----------------------------------+
|            | *.smtp.github.com                 |
+            +-----------------------------------+
|            | *.smtp.github.com                 |
|            | smtp.github.com                   |
+            +-----------------------------------+
|            | skyline.github.com                |
+            +-----------------------------------+
|            | skyline.github.com                |
|            | www.skyline.github.com            |
+            +-----------------------------------+
|            | api.security.github.com           |
+            +-----------------------------------+
|            | api.security.github.com           |
|            | www.api.security.github.com       |
+            +-----------------------------------+
|            | f.cloud.github.com                |
+            +-----------------------------------+
|            | *.registry.github.com             |
+            +-----------------------------------+
|            | *.registry.github.com             |
|            | registry.github.com               |
+            +-----------------------------------+
|            | help.github.com                   |
+            +-----------------------------------+
|            | docs.github.com                   |
|            | help.github.com                   |
+            +-----------------------------------+
|            | support.enterprise.github.com     |
+            +-----------------------------------+
|            | support.enterprise.github.com     |
|            | www.support.enterprise.github.com |
+            +-----------------------------------+
|            | jira.github.com                   |
+            +-----------------------------------+
|            | jira.github.com                   |
|            | www.jira.github.com               |
+            +-----------------------------------+
|            | fast.github.com                   |
+            +-----------------------------------+
|            | r2.shared.global.fastly.net       |
+            +-----------------------------------+
|            | dns-vetting1i.map.fastly.net      |
+            +-----------------------------------+
|            | mac-installer.github.com          |
+            +-----------------------------------+
|            | vscode-auth.github.com            |
+            +-----------------------------------+
|            | classroom.github.com              |
+            +-----------------------------------+
|            | styleguide.github.com             |
+            +-----------------------------------+
|            | render.github.com                 |
+            +-----------------------------------+
|            | render.github.com                 |
|            | render-lab.github.com             |
+            +-----------------------------------+
|            | www.github.com                    |
+            +-----------------------------------+
|            | *.codespaces-ppe.github.com       |
+            +-----------------------------------+
|            | *.codespaces-ppe.github.com       |
|            | codespaces-ppe.github.com         |
+            +-----------------------------------+
|            | *.codespaces.github.com           |
+            +-----------------------------------+
|            | *.codespaces.github.com           |
|            | codespaces.github.com             |
+            +-----------------------------------+
|            | *.codespaces-dev.github.com       |
+            +-----------------------------------+
|            | *.codespaces-dev.github.com       |
|            | codespaces-dev.github.com         |
+            +-----------------------------------+
|            | lab.github.com                    |
+            +-----------------------------------+
|            | *.workspaces-dev.github.com       |
+            +-----------------------------------+
|            | *.workspaces-dev.github.com       |
|            | workspaces-dev.github.com         |
+            +-----------------------------------+
|            | *.workspaces.github.com           |
+            +-----------------------------------+
|            | *.workspaces.github.com           |
|            | workspaces.github.com             |
+            +-----------------------------------+
|            | *.workspaces-ppe.github.com       |
+            +-----------------------------------+
|            | *.workspaces-ppe.github.com       |
|            | workspaces-ppe.github.com         |
+            +-----------------------------------+
|            | graphql-stage.github.com          |
+            +-----------------------------------+
|            | graphql-stage.github.com          |
|            | www.graphql-stage.github.com      |
+            +-----------------------------------+
|            | graphql.github.com                |
+            +-----------------------------------+
|            | graphql.github.com                |
|            | www.graphql.github.com            |
+            +-----------------------------------+
|            | maintainers.github.com            |
+            +-----------------------------------+
|            | maintainers.github.com            |
|            | www.maintainers.github.com        |
+            +-----------------------------------+
|            | *.a.ssl.fastly.net                |
+            +-----------------------------------+
|            | f.cloud.github.com                |
|            | raw.github.com                    |
+            +-----------------------------------+
|            | import.github.com                 |
+            +-----------------------------------+
|            | import.github.com                 |
|            | porter.github.com                 |
+            +-----------------------------------+
|            | pkg.github.com                    |
+            +-----------------------------------+
|            | lab-sandbox.github.com            |
+            +-----------------------------------+
|            | *.staging-lab.github.com          |
+            +-----------------------------------+
|            | *.staging-lab.github.com          |
|            | staging-lab.github.com            |
+            +-----------------------------------+
|            | brandguide.github.com             |
+            +-----------------------------------+
|            | talks.github.com                  |
+            +-----------------------------------+
|            | learn.github.com                  |
+            +-----------------------------------+
|            | status.github.com                 |
+            +-----------------------------------+
|            | octostatus-production.github.com  |
|            | status.github.com                 |
+            +-----------------------------------+
|            | slack.github.com                  |
+            +-----------------------------------+
|            | cla.github.com                    |
+            +-----------------------------------+
|            | GitHub,                           |
|            | Inc.                              |
+            +-----------------------------------+
|            | support@github.com                |
+            +-----------------------------------+
|            | cloud.github.com                  |
+            +-----------------------------------+
|            | enterprise.github.com             |
+            +-----------------------------------+
|            | education.github.com              |
+            +-----------------------------------+
|            | education.github.com              |
|            | edu.github.com                    |
+            +-----------------------------------+
|            | *.branch.github.com               |
+            +-----------------------------------+
|            | *.branch.github.com               |
|            | branch.github.com                 |
+            +-----------------------------------+
|            | jobs.github.com                   |
+            +-----------------------------------+
|            | community.github.com              |
+            +-----------------------------------+
|            | *.hq.github.com                   |
+            +-----------------------------------+
|            | *.hq.github.com                   |
|            | hq.github.com                     |
+            +-----------------------------------+
|            | *.stg.github.com                  |
+            +-----------------------------------+
|            | *.stg.github.com                  |
|            | stg.github.com                    |
+            +-----------------------------------+
|            | schrauger.github.io               |
+            +-----------------------------------+
|            | github.com                        |
|            | schrauger.github.com              |
+            +-----------------------------------+
|            | a.ssl.fastly.net                  |
+            +-----------------------------------+
|            | offer.github.com                  |
+            +-----------------------------------+
|            | atom-installer.github.com         |
+            +-----------------------------------+
|            | *.rs.github.com                   |
+            +-----------------------------------+
|            | *.rs.github.com                   |
|            | rs.github.com                     |
+            +-----------------------------------+
|            | dodgeball.github.com              |
+            +-----------------------------------+
|            | helpnext.github.com               |
+            +-----------------------------------+
|            | *.github.com                      |
|            | www.github.com                    |
+            +-----------------------------------+
|            | central.github.com                |
+            +-----------------------------------+
|            | *.id.github.com                   |
+            +-----------------------------------+
|            | *.id.github.com                   |
|            | id.github.com                     |
+            +-----------------------------------+
|            | camo.github.com                   |
+------------+-----------------------------------+

🐳 You can pull docker image in: https://hub.docker.com/r/ductn4/loo

⭐️ Star the project on GitHub if you liked this tool

👉 https://github.com/ductnn/lookup 👈

🎉 Thank you so much 🎉
```

## Download

For linux x64: [download](./bin/loo_linux)
For MacOS x64: [download](./bin/loo_darwin)

## License
The MIT License (MIT). Please see [LICENSE](license) for more information.
