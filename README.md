# tfsubst

tfstate substitution tool like `envsubst`.

## Usage

`tfsubst` substitutes all variables in the file with the values from the tfstate.

```bash
$ tfsubst -s <tfstate-file> -i <input-file>
```

### Example

If you have a template file below

```yaml
# input.yaml
value: {{ tfstate "aws_ssm_parameter.your_parameter.value" }}
```

and your tfstate is in `s3://yourbucket/terraform.tfstate`

then the output file will be

```bash
$ tfsubst -s s3://yourbucket/terraform.tfstate -i input.yaml -o output.yaml
```

```yaml
# output.yaml
value: your_parameter_value
```

## Installation

### From binary

Download the binary from [GitHub Releases](https://github.com/RossyWhite/tfsubst/releases) and drop it in your `$PATH`.

### With Go

```shell
$ go install github.com/RossyWhite/tfsubst
```

### Docker

```shell
$ docker pull ghcr.io/rossywhite/tfsubst:latest
```
