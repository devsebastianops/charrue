![charrue](./docs/logo/cover.png)


**charrue** is a tool that generates Terraform configuration files from a higher-level config and template directory. Designed to abstract away the limitations of HCL, charrue lets you write infrastructure logic using real programming constructs — without giving up the power of Terraform.

> **charrue** is a French word for *plow*.
> And just like a plow, it prepares the ground for something to grow.


## :seedling: Why charrue?

Terraform's HCL is declarative, but painfully limited when it comes to reusability, logic, and abstractions. charrue introduces a programmable layer between you and HCL, allowing you to:

- Write DRY, logical, and testable infrastructure configs
- Use real `if`/`else`, loops, functions, and composition
- Maintain Terraform compatibility (charrue emits `.tf` files)

## :rocket: Quickstart

```bash
# Generate terraform files from the current directory
charrue build

# Alternatively you can specify other directories.
# For more information, check out the help command, or read the documentation.

# Change directory to the output directory
cd dist

# Next steps: Run terraform commands
terraform init
terraform plan --out planfile
terraform apply planfile
```

Check out the [documentation](/docs/README.md).

## :wrench: Installation


### Go install

```bash
go install github.com/devsebastianops/charrue@latest
```

### Install from source

```bash
git clone https://github.com/devsebastianops/charrue.git
cd charrue
make install
```

## :pencil: Roadmap

- [x] Rendering of Terraform files using go templates
- [x] Data from YAML files
- [ ] Support for other input data formats (e.g. JSON, CSV)
- [ ] Support for other template engines (e.g. Jinja2, Mustache)
- [ ] Configuration file for charrue
- [ ] Package managers support (e.g. `brew`, `scoop`, etc.)
- [ ] Docker support

## :ledger: Licensee
charrue is licensed under the [Apache 2 License](LICENSE). Feel free to use, modify, and distribute it under the terms of this license. Contributions are welcome!

