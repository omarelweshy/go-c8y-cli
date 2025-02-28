# go-c8y-cli

![build](https://github.com/reubenmiller/go-c8y-cli/workflows/build/badge.svg)

<p align="center">
    <img width="1000" src="demo.svg">
</p>


Cumulocity IoT Command Line Tool

Supported on

* Linux (amd64, x86, armv5->7)
* MacOS (amd64, arm64)
* Windows (amd64, x86, arm64 (binary only))

## Installation

See the following installation instructions

* [Shell](https://goc8ycli.netlify.app/docs/installation/shell-installation)
* [Docker](https://goc8ycli.netlify.app/docs/installation/docker-installation)
* [PowerShell](https://goc8ycli.netlify.app/docs/installation/powershell-installation)


## Documentation

See the [documentation website](https://goc8ycli.netlify.app/) for instructions on how to install and use it.

## Contributing

1. Fork the project, then clone it

    ```sh
    git clone https://github.com/reubenmiller/go-c8y-cli.git
    ```

2. Optional: If you have existing .cumulocity sessions folder, then you can copy the files into the local directory so that they are available for use during development

    ```sh
    cd go-c8y-cli

    # bash/zsh
    mkdir -p .cumulocity
    cp -R ~/.cumulocity/ .cumulocity/
    ```

3. Open the project in Microsoft VS Code (using Dev Containers - this requires Docker!)

    ```sh
    code go-c8y-cli

    # When prompted, build and open the dev container
    ```

4. Run initial setup tasks so that you can run c8y inside the dev container

    ```sh
    task init-setup
    ```

5. Add or edit a command specification (`.yaml` file) in `api/spec/yaml/`. The specifications are used to auto generate the go code

6. Run the code generation and build the go binary

    ```sh
    task generate build-snapshot-single

    # reload your shell
    zsh
    ```

7. Try out the newly built binary (it should already be added to your)

    **Shell**

    ```bash
    c8y currentuser get
    ```

    **PowerShell**

    ```powershell
    task generate build-powershell

    pwsh
    Import-Module ./tools/PSc8y/dist/PSc8y -Force
    Get-CurrentUser
    ```

### Building the documentation

1. Update the auto generated cli docs (if you have changed something)

    ```sh
    task docs
    ```

2. Launch the documentation preview

    ```sh
    task gh-pages
    ```

3. View the documentation in the [browser](http:/localhost:3000)


## Tests

### Pre-requisites

1. Build the latest version and update auto generated tests

    ```sh
    task build
    task generate-cli-tests
    ```

1. Set the c8y session that you want to use for the tests

    ```sh
    set-session
    ```

### Run test on example code

The examples included in the API specification can be validated by running the follow make task.

```sh
task test-cli
```

### Run powershell tests

```sh
task test-powershell
```
