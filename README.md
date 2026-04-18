# GitHub User Activity CLI

A simple Command Line Interface (CLI) application built with Go to fetch and display recent activity of a GitHub user.

This project is a solution to the [GitHub User Activity](https://roadmap.sh/projects/github-user-activity) challenge from roadmap.sh.

## Features

- Fetch recent public events of a GitHub user using the GitHub API.
- Validate GitHub username format before making requests.
- Display activity types such as:
  - Pushing commits.
  - Opening/closing issues.
  - Starring repositories.
  - Forking repositories.
  - Creating branches or tags.
- Simple and clean CLI output.

## Prerequisites

- [Go](https://golang.org/doc/install).

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/rizkyfauziilmi/github-user-activity-go.git
   cd github-user-activity-go