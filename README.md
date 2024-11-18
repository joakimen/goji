# goji

[![CI](https://github.com/joakimen/goji/actions/workflows/ci.yml/badge.svg)](https://github.com/joakimen/goji/actions/workflows/ci.yml)

CLI for listing various Jira issue types.

## CLI Surface

- `goji`
  - `auth`
    - `login`: Add Jira credentials to the system keyring
    - `show`: Show stored credentials
    - `status`: Show authentication status
  - `epic`
    - `list [flags]`
      - `-p, --project`: Project key
      - `-a, --all`: Return both resolved and unresolved epics
      - `-j, --json`: Output as JSON
      - `-m, --mine`: Return only epics assigned to the current user
  - `issue`
    - `list [flags]`
      - `-p, --project`: Project key
      - `-a, --all`: Return both resolved and unresolved issues
      - `-j, --json`: Output as JSON
      - `-m, --mine`: Return only issues assigned to the current user
      - `-l, --limit`: Maximum number of issues to return (default: 50)

## Authentication

Jira credentials are stored securely in the system keyring via the `auth` subcommand.
