# goji

CLI for listing various Jira issue types.

## CLI Surface

- `goji`
    - `auth`
        - `login`: Store Jira credentials in the system keyring
        - `show`: Print stored Jira credentials
        - `clear`: TODO - Remove stored Jira credentials
    - `epic`
        - `list`

    later..
    - `issue`
        - `list`
    - `bug`
        - `list`

## Authentication

Jira credentials are stored securely in the system keyring via the `auth` subcommand.
