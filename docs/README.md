# Documentation

## `build`

### Values file
Per default charrue will look for a `values.yaml` file in the current directory. You can override this by using the `--values` or `--V` flag.

```bash
charrue build --values ~/my-values.yaml
```

### Template path
Per default charrue will find all templates in the current directory. Currently charrue is capable of `.tmpl` file extension and go templates. You can override this by using the `--template-path` or `-t` flag.

```bash
charrue build --template-path ~/my-templates
```

### Output path
Per default charrue will output the generated terraform files in the `dist` directory. You can override this by using the `--output-path` or `-o` flag.

```bash
charrue build --output-path ~/my-output
```

### Recursive behavior
Normally charrue will not look for templates in subdirectories. You can override this by using the `--recursive` or `-r` flag.

```bash
charrue build --recursive
```

### Terraform format
Per default charrue will format the generated terraform files using `terraform fmt`. You can override this by using the `--no-format` flag.

```bash
charrue build --no-format
```

