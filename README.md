# Docker Datacenter Compliance Controls [![CircleCI](https://circleci.com/gh/docker/ddc-opencontrol/tree/master.svg?style=svg&circle-token=daeaf5acd7ac08000ea727cbf8ec9baa8ded8da4)](https://circleci.com/gh/docker/ddc-opencontrol/tree/master) [![codecov](https://codecov.io/gh/docker/ddc-opencontrol/branch/master/graph/badge.svg?token=WiRPQcno3c)](https://codecov.io/gh/docker/ddc-opencontrol)

This repository contains compliance information for Docker Datacenter as it pertains to NIST-800-53 Rev 4 security controls and the FedRAMP Moderate and High baselines. This data adheres to the [OpenControl](http://open-control.org/) schema for building compliance documentation and can be used as part of your own ATO review process. The documentation generated from this content can be used to ATO Docker Datacenter in both on-premises/private cloud infrastructure and in public cloud providers.

> This content is provided for informational purposes only and has not been vetted by any third-party security assessors. You are solely responsible for developing, implementing, and managing your applications and/or subscriptions running on your own platform in compliance with applicable laws, regulations, and contractual obligations. The documentation is provided "as-is" and without any warranty of any kind, whether express, implied or statutory, and Docker, Inc. expressly disclaims all warranties for non-infringement, merchantability or fitness for a particular purpose.

Pre-built System Security Plan (SSP) templates for running Docker Datacenter on various cloud providers as indicated in the table below can be obtained by contacting [compliance@docker.com](mailto:compliance@docker.com). Note that even if a pre-built template for DDC is not available for your chosen cloud provider, you can still use the OpenControl-formatted content in this repository to generate your own SSP templates. Much of the content in this repository is identical to that which is provided in the pre-built templates.

|Provider|Format|Baselines|Status|
|--------|------|---------|------|
|[Microsoft Azure Government](https://azure.microsoft.com/en-us/overview/clouds/government/)|[Azure Blueprint](https://blogs.msdn.microsoft.com/azuregov/2016/10/12/azure-blueprint-architecting-secure-solutions-just-got-easier/) (.docx)|Moderate<br>High<br>DoD L4<br>DoD L5|Available<br>Coming Soon<br>Coming Soon<br>Coming Soon|
|[AWS GovCloud](https://aws.amazon.com/govcloud-us/)|TBD|Moderate|Coming soon|

## Usage

The control guidance for Docker Datacenter is separated in to the following components:

|Component Name|Folder|Version|
|--------------|------|-------|
|Commercially Supported (CS) Docker Engine|[`CSEngine/`](https://github.com/docker/ddc-opencontrol/tree/master/CSEngine)|1.12.3-cs4|
|Docker Trusted Registry (DTR)|[`DTR/`](https://github.com/docker/ddc-opencontrol/tree/master/DTR)|2.1.1|
|Universal Control Plane (UCP)|[`UCP/`](https://github.com/docker/ddc-opencontrol/tree/master/UCP)|2.0.1|
|Authentication and Authorization Service|[`eNZi/`](https://github.com/docker/ddc-opencontrol/tree/master/eNZi)|2.0.1|

> Both the UCP and DTR components leverage the eNZi authentication and authorization Service component for authentication and authorization across an entire Docker Datacenter cluster.

A `component.yaml` file resides in each component's subdirectory. Updates to the security narratives and content are made to these `component.yaml` files.

In order to generate the documentation appropriate to your system, refer to the Compliance Masonry [usage instructions](https://github.com/opencontrol/compliance-masonry). The `examples/ddc-compliance` directory contains an example use of these components.

## Developing

Refer to the [Contributing Guide](https://github.com/docker/ddc-opencontrol/blob/master/CONTRIBUTING.md) for instructions on contributing to this project.

### Component Validation

The OpenControl schema is defined by the [Kwalify](http://www.kuwata-lab.com/kwalify/) schema validator and YAML parser. Each Docker Datacenter component definition is tested against this schema using the [PyKwalify](https://github.com/Grokzen/pykwalify) Python port of the Kwalify specification. This repository contains a Dockerfile for running the component tests within a container.

```sh
docker build -t docker/ddc-opencontrol .
docker run docker/ddc-opencontrol
```

### Natural Language Processing [Experimental]

Thorough documentation of the relevant security controls for each of the DDC components is a critical aspect of this project. It's imperative that not only is each control satisfied, but that the contents of the actual narratives match that which is defined by NIST 800-53. As such, this project includes experimental support for key phrase text analysis backed by [Microsoft Cognitive Services](https://www.microsoft.com/cognitive-services).

The [`nlp/`](https://github.com/docker/ddc-opencontrol/tree/master/nlp) directory contains a command-line service written in Go that parses each component control's narratives and submits them to the [Text Analytics API](https://www.microsoft.com/cognitive-services/en-us/text-analytics-api) for detection of key phrases (e.g. "access control", "authentication", etc). The key phrases are then checked against the key phrases that represent each of the NIST 800-53 definitions to ensure that the content indeed matches. The match process is currently quite basic. A successful match occurs when a component's narrative includes one or more of the key phrases that are also in the list of key phrases representative of the NIST definition itself. You can think of this as a form of automated proofreading.

Ultimately, this functionality is best served as a compliance-masonry [plugin](https://github.com/opencontrol/compliance-masonry/tree/master/exampleplugin) developed in a separate repository instead of as a standalone tool. Contributions welcome!

The potential use cases for natural language processing in documentation efforts are pretty wide-ranging. As such, our goal with this example is to open the door to new and exciting ways to build security and compliance documentation.
