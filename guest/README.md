<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
***
***
***
*** To avoid retyping too much info. Do a search and replace for the following:
*** github_username, repo_name, twitter_handle, email, project_title, project_description
-->

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/patricksferraz/guest">
    <img src="img/logo.png" alt="Logo" width="100" height="100">
  </a>

  <h3 align="center">Guest</h3>

  <p align="center">
    Microservice for time recording
    <br />
    <a href="https://github.com/patricksferraz/guest"><strong>Explore the docs »</strong></a>
    <!-- <br />
    <br />
    <a href="https://github.com/patricksferraz/guest">View Demo</a>
    ·
    <a href="https://github.com/patricksferraz/guest">Report Bug</a>
    ·
    <a href="https://github.com/patricksferraz/guest">Request Feature</a>-->
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <!-- <li><a href="#usage">Usage</a></li> -->
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <!-- <li><a href="#license">License</a></li> -->
    <li><a href="#contact">Contact</a></li>
    <!-- <li><a href="#acknowledgements">Acknowledgements</a></li> -->
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

Guest is a microservice for pinned soluction providing in the application layer the communication by REST and gRPC.

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->
<!--
Here's a blank template to get started:
**To avoid retyping too much info. Do a search and replace with your text editor for the following:**
`github_username`, `repo_name`, `twitter_handle`, `email`, `project_title`, `project_description` -->

### Built With

- [Go Lang](https://golang.org/)
- List all: `go list -m all`

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- Hiring a kubernetes cluster:
  - [AWS](https://aws.amazon.com/pt/eks/?whats-new-cards.sort-by=item.additionalFields.postDateTime&whats-new-cards.sort-order=desc&eks-blogs.sort-by=item.additionalFields.createdDate&eks-blogs.sort-order=desc)
  - [Azure](https://azure.microsoft.com/pt-br/services/kubernetes-service/)
  - [GCP](https://cloud.google.com/kubernetes-engine)

- [Kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)

- Create a secrets: see "_k8s/instructions.md_"

### Deploy

- `kubectl apply -f ./k8s`

<!-- USAGE EXAMPLES -->
<!-- ## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_ -->

<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/patricksferraz/guest/issues) for a list of proposed features (and known issues).

<!-- CONTRIBUTING -->
## Contributing

Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

**Prerequisites**:

- Golang

  ```sh
  wget https://golang.org/dl/go1.17.7.linux-amd64.tar.gz
  rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz
  export PATH=$PATH:/usr/local/go/bin
  ```

- Docker and docker-compose

  ```sh
  sudo apt-get install docker docker-compose docker.io -y
  ```

- Environment: see "_.env.example_"

**Installation**:

1. Clone the repo

   ```sh
   git clone https://github.com/patricksferraz/guest.git
   ```

2. Run

   ```sh
   make up
   ```

3. Test

   ```sh
   make gtest
   ```

**Installation in local kubernetes**:

1. Install [k3d](https://k3d.io/), [Kind](https://kind.sigs.k8s.io/) or similar
2. Install [Kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) and [Helm](https://helm.sh/)
3. Follow the steps of [Getting Started](#getting-started)
    - Connect to cluster and run:

      `kubectl apply -f k8s/`
<!-- LICENSE -->
<!-- ## License -->

<!-- Distributed under the MIT License. See `LICENSE` for more information. -->

<!-- CONTACT -->
## Contact

Coding4u - contato@coding4u.com.br - [website](http://coding4u.com.br)

Project Link: [guest](https://github.com/patricksferraz/guest)

<!-- ACKNOWLEDGEMENTS -->
<!-- ## Acknowledgements

* []()
* []()
* []() -->
