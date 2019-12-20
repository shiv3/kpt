// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package readme

var READMEShort = `  Git based configuration package manager`
var READMELong = `
### Synopsis

  Git based configuration package manager.

**Packages are composed of Resource configuration** (rather than DSLs, templates, etc), but may
also contain supplemental non-Resource artifacts (e.g. README.md, arbitrary other files).

  Resource configuration is a collection of Kubernetes style objects (yaml or json)
  stored in files:

        # dir/deployment.yaml
        apiVersion: apps/v1
        kind: Deployment
        metadata:
          name: petclinic-frontend
        ...

  Packaging configuration rather than Templates or DSLs provides a number of desirable properties
  such as:

  - it clearly **represents the intended state** of the infrastructure -- no for loops, http calls,
    etc

  - it **works with Kubernetes project based tools**

  - it lends itself to the **development of new / custom tools**
    - new tools can be developed read and modify the package contents based on the Resource schema.
    - validation and linting tools (e.g. ` + "`" + `kubeval` + "`" + `)
    - parsing and modifying via the cli (e.g. ` + "`" + `kustomize config set` + "`" + `)
    - parsing and modifying declaratively through meta Resources
      (e.g. ` + "`" + `kustomize` + "`" + `, ` + "`" + `kustomize config run` + "`" + `)

  - tools may be written in **any language or framework**
    - tools just manipulate yaml / json directly, rather than manipulating Templates or DSLs
    - can use Kubernetes language libraries and openapi schema

**Every existing git subdirectory containing Resource configuration** may be used as a ` + "`" + `kpt` + "`" + `
package.

  Requirement for packages: -- they are **stored in a git and are directories containing Resource
  configuration**.
  Notably, the upstream [https://github.com/kubernetes/examples/staging/cockroachdb] qualifies
  as a package:

    # fetch the examples cockroachdb directory as a package
    kpt get https://github.com/kubernetes/examples/staging/cockroachdb my-cockroachdb

**Packages use git references for versioning**.

  Package consumers may target a version using git tags, branches, commits etc.  Package
  publishers are encouraged to adopt semantic versioning.

    # fetch the examples cockroachdb directory as a package
    kpt get https://github.com/GoogleContainerTools/kpt/examples/cockroachdb@v0.1.0 my-cockroachdb

**Packages may be customized through either in place modifications or through expansion**.

  It is possible to directly modify a fetched package.  Updates from upstream may be merged
  into the local package.  Some packages may expose *field setters* used by kustomize to change
  specific fields.

    export KUSTOMIZE_ENABLE_ALPHA_COMMANDS=true # enable alpha kustomize commands

    kpt get https://github.com/GoogleContainerTools/kpt/examples/cockroachdb my-cockroachdb
    kustomize config set my-cockroachdb/ replicas 5

  It is also possible to indirectly customize the packages by applying modifications to expanded
  Resources -- e.g. via Kustomize:

    kpt get https://github.com/GoogleContainerTools/kpt/examples/cockroachdb my-cockroachdb
    # create kustomizations
    ...
    kustomize build my-cockroachdb/

**The same kpt package may be fetched multiple times** to separate locations in order to **create
separate instances**.

  Each instance may be modified and updated independently of the others.

    export KUSTOMIZE_ENABLE_ALPHA_COMMANDS=true # enable alpha kustomize commands

    # fetch an instance of a java package
    kpt get https://github.com/GoogleContainerTools/kpt/examples/java my-java-1
    kustomize config set my-java-1/ image gcr.io/example/my-java-1:v3.0.0

    # fetch a second instance of a java package
    kpt get https://github.com/GoogleContainerTools/kpt/examples/java my-java-2
    kustomize config set my-java-2/ image gcr.io/example/my-java-2:v2.0.0

**Packages may pull in updates** from the upstream package in git.

 Specify the target version to update to, and an (optional) update strategy for how to apply the
 upstream changes -- strategies may merge Resources by field, merge files by line number,
 replace files, or fail on local changes.

    export KUSTOMIZE_ENABLE_ALPHA_COMMANDS=true # enable alpha kustomize commands

    kpt get https://github.com/GoogleContainerTools/kpt/examples/cockroachdb my-cockroachdb
    kustomize config set my-cockroachdb/ replicas 5
    kpt update my-cockroachdb@v1.0.1 --strategy=resource-merge

#### Layering and Composition

Each Kubernetes Resource has a number of different fields.  In many cases **different field
values may be defined by different teams** -- e.g. a platform team may want to add a sidecar,
an SRE team may want to set replicas, cpu, memory, a dev team may set environment variables
or images.

Using a Resource-centric approach to packaging is more amenable to unifying opinions of multiple
teams by merging Resources.  When stored in ` + "`" + `yaml` + "`" + `, Resource fields may be annotated with
the last setter of the field.

Example of a Resource annotated with field origins.

    # Deployment unifying the opinions of platform, petclinic-dev and app-sre teams
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: petclinic-frontend
      namespace: petclinic-prod # {"setBy":"app-sre"}
      labels:
        app: petclinic-frontend # {"setBy":"petclinic-dev"}
        env: prod # {"setBy":"app-sre"}
    spec:
      replicas: 3 # {"setBy":"app-sre"}
      selector:
        matchLabels:
          app: petclinic-frontend # {"setBy":"petclinic-dev"}
          env: prod # {"setBy":"app-sre"}
      template:
        metadata:
          labels:
            app: petclinic-frontend # {"setBy":"petclinic-dev"}
            env: prod # {"setBy":"app-sre"}
      spec:
          containers:
          - name: petclinic-frontend
            image: gcr.io/petclinic/frontend:1.7.9 # {"setBy":"app-sre"}
            args:
            - java # {"setBy":"platform"}
            - -XX:+UnlockExperimentalVMOptions # {"setBy":"platform"}
            - -XX:+UseCGroupMemoryLimitForHeap # {"setBy":"platform","description":"dynamically determine heap size"}
            ports:
            - name: http
              containerPort: 80 # {"setBy":"platform"}

#### Templates and DSLs

Note: If the use of Templates or DSLs is strongly desired, they may be used to produce
` + "`" + `kpt` + "`" + ` packages by fully expanding them into Resource configuration.

The artifacts used to generated Resource configuration may be included in the package as
supplements.

#### Flags

  --stack-trace

    Print a stack trace on an error.  For debugging code.

#### Env Vars

  COBRA_SILENCE_USAGE

    Set to true to silence printing the usage on error

  COBRA_STACK_TRACE_ON_ERRORS

    Set to true to print a stack trace on an error

  KPT_NO_PAGER_HELP

    Set to true to print the help to the console directly instead of through
    a pager (e.g. less)`