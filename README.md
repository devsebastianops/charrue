# Charrue

**Charrue** is a tool that generates Terraform configuration files from a higher-level config and template directory. Designed to abstract away the limitations of HCL, Charrue lets you write infrastructure logic using real programming constructs — without giving up the power of Terraform.

> **Charrue** is a French word for *plow*.
> And just like a plow, it prepares the ground for something to grow.


## :seedling: Why Charrue?

Terraform's HCL is declarative, but painfully limited when it comes to reusability, logic, and abstractions. Charrue introduces a programmable layer between you and HCL, allowing you to:

- Write DRY, logical, and testable infrastructure configs
- Use real `if`/`else`, loops, functions, and composition
- Maintain Terraform compatibility (Charrue emits `.tf` files)

## :rocket: Quickstart

```bash
# Generate .tf files from your config and templates
charrue render config.yaml templates/ > main.tf

# Optionally run Terraform
terraform init
terraform apply
```

## :pencil: Features
- High-level config input (YAML or JSON)
- Template-based Terraform generation
- Optional terraform validate and formatting
- CLI-first, Git-friendly, CI/CD-ready

## :ledger: Licensee
Charrue is licensed under the [Apache 2 License](LICENSE). Feel free to use, modify, and distribute it under the terms of this license. Contributions are welcome!

