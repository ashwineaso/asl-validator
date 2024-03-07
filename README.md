# github.com/ashwineaso/asl-validator
A simple Amazon States Language parser and validator based on JSON schemas. 

This validator makes it possible for you to locally validate your state machine definition without actually creating the state machine

### Installation
```bash
go install github.com/ashwineaso/asl-validator
```

### CLI Usage
```
> go-asl-parser --help
NAME:
   go-asl-validator - validate the ASL file

USAGE:
   go-asl-validator [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --json-path FILE  json path to the ASL file FILE
   --help, -h        show help
```

Currently only files containing JSON schema can be validated.


### References
- [ASL specifications](https://states-language.net/spec.html)
- [ASL Documentation on AWS](http://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html)


### License

See [LICENSE](./LICENSE).