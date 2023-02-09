# Codecharacter 2023 LSP Server

---

### Requirements

- [Go](https://go.dev/)
- [golangci](https://golangci-lint.run/usage/install/)
- [Docker](https://www.docker.com/)

### Setup

- Configure .vscode/settings.json
  ```
  {
      "go.lintTool":"golangci-lint",
      "go.lintFlags": [
      "--fast"
      ],
      "go.lintOnSave": "package",
      "go.formatTool": "goimports",
      "go.useLanguageServer": true,
      "[go]": {
          "editor.formatOnSave": true,
          "editor.codeActionsOnSave": {
              "source.organizeImports": true
          }
      },
      "go.docsTool": "gogetdoc"
  }
  ```
- Create .env file
  ```sh
  cp .env.example .env
  ```
- Init Submodules
  ```sh
  git submodule update --init
  ```
- Enable githooks
  ```sh
  git config core.hooksPath .githooks
  ```

### Run

- #### On Docker
  ```sh
  docker-compose up
  ```
