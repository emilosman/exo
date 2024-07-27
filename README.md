# exo-cli
_"The best way to learn is to rewrite something that exists"_

- CLI helper for my [exocortex](https://en.wikipedia.org/wiki/Exobrain)

## Bonzai references
This CLI is built with [rwxrob](https://github.com/rwxrob)'s Bonzai template.  
Here are some useful links that helped me understand how to use it:

- [bonzai repo](https://github.com/rwxrob/bonzai)
- [z repo](https://github.com/rwxrob/z)
- [z repo main.go](https://github.com/rwxrob/z/blob/main/main.go)
- [example template](https://github.com/rwxrob/bonzai-example)
- [bonzai help](https://github.com/rwxrob/help)
- [Bonzai and how to create a personal CLI to rule them all](https://dev.to/cherryramatis/bonzai-and-how-to-create-a-personal-cli-to-rule-them-all-1bnl)

## Todo
- [ ] pages functions
    - [ ] search, suggest page when not found if similar exists: `exo page ...`
    - [ ] alias command `p = pages`
    - [x] list all pages: `exo pages`
- [ ] config file
    - [ ] set path for _exo_ folder
- [ ] hyperlinks, references
- [ ] `exo init`
- [ ] other functionality
    - [ ] github sync
    - [ ] go build
- [ ] exo secret vault

## Done
- [x] daily functions
    - [x] create _today_ if it does not exist
    - [x] openDay function: `exo day YYYYMMDD`
    - [x] yesterday: `exo yesterday`
    - [x] today: `exo today`
