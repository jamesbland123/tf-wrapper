# Terraform wrapper made with Golang

This is a wrapper program used to run terraform pipelines with a single command. 
It's primarily meant as a working example to modify and expand on.  Versus other examples
you might find online I tried implementations that are not as trivial and 
something that might actually get used.  

**Features**
- Cobra for command line args and flags
- Viper for config file to record common commands
- Configurable to run a pipeline of terraform configurations/directories
- Support for different environments: dev/integration/prod...
- Docker container support with image name arg
- Plus more...

After compiling run
```bash
$ tf-wrapper -h  # for help

$ tf-wrapper apply -e dev -i abc123  # example apply pipeline
```
