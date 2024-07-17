<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">SIMPLE-L4-LB</h1>
</p>
<p align="center">
    <em>Directing traffic, maximizing performance.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/buildflash/simple-l4-lb?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/buildflash/simple-l4-lb?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/buildflash/simple-l4-lb?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/buildflash/simple-l4-lb?style=default&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>

<br><!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary><br>

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ Getting Started](#-getting-started)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Tests](#-tests)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)
</details>
<hr>

##  Overview

Simple-l4-lb simplifies load balancing through a CLI, offering Round Robin and Consistent Hashing strategies for managing backend servers. The projects main functionality lies in routing incoming requests to appropriate servers based on chosen strategies dynamically. With the ability to add and adjust backends on the fly, simple-l4-lb provides a valuable solution for network service optimization and flexibility, enhancing overall system performance.

---



##  Modules


| File                                                                              | Summary                                                                                                                                                                                                                                                                             |
| ---                                                                               | ---                                                                                                                                                                                                                                                                                 |
| [main.go](https://github.com/buildflash/simple-l4-lb/blob/master/main.go)         | Executes a command-line interface (CLI) for managing load balancer settings. Commands include changing strategies, adding backend servers, and testing topologies. Handles input and output, sets logging, and initializes the load balancer.                                       |
| [strategy.go](https://github.com/buildflash/simple-l4-lb/blob/master/strategy.go) | Implements strategies for load balancing backends in a network service. RR strategy cycles through backends for each request. Consistent Hashing strategy maps requests to specific backends using a hash ring. Provides methods to manage, choose, and display backend topologies. |
| [lb.go](https://github.com/buildflash/simple-l4-lb/blob/master/lb.go)             | Implements load balancing logic with backend management. It handles incoming requests, proxies them to proper backends based on set strategy, and can dynamically change strategies or add backends at runtime.                                                                     |

</details>

---

##  Getting Started

**System Requirements:**

* **Go**: `version 1.21.0`

###  Installation

<h4>From <code>source</code></h4>

> 1. Clone the simple-l4-lb repository:
>
> ```console
> $ git clone https://github.com/buildflash/simple-l4-lb
> ```
>
> 2. Change to the project directory:
> ```console
> $ cd simple-l4-lb
> ```
>
> 3. Install the dependencies:
> ```console
> $ go build -o myapp
> ```

###  Usage

<h4>From <code>source</code></h4>

> Run simple-l4-lb using the command below:
> ```console
> $ ./myapp
> ```


##  Project Roadmap

- [ ] `► Add Graceful Exit support`
- [ ] `► Add more load balancing strategies`
- [ ] `► Add support for receiving config events from redis pubsub`
- [ ] `► Add health checks for backends`

---

##  Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Report Issues](https://github.com/buildflash/simple-l4-lb/issues)**: Submit bugs found or log feature requests for the `simple-l4-lb` project.
- **[Submit Pull Requests](https://github.com/buildflash/simple-l4-lb/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/buildflash/simple-l4-lb/discussions)**: Share your insights, provide feedback, or ask questions.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/buildflash/simple-l4-lb
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="center">
   <a href="https://github.com{/buildflash/simple-l4-lb/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=buildflash/simple-l4-lb">
   </a>
</p>
</details>

---

##  License

This project is protected under the Apache License. For more details, refer to the [LICENSE](LICENSE) file.

---

##  Acknowledgments

- Readme made with [readme-ai](https://github.com/eli64s/readme-ai)
- Inspired by [Arpit Bhayani](https://arpitbhayani.me/)

[**Return**](#-overview)

---
