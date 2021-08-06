# makefile-completion-go`

## Install

1. $ `go get -v -u github.com/tomlla/makefile-completion-go`
2. Paste this bash script to your ~/.bashrc

```bash
function _makefile_targets_01 {
    local curr_arg;
    local targets;
    targets=''
    if [[ -e "$(pwd)/Makefile" ]]; then
        targets=$(makefile-completion-go)
    fi
    curr_arg=${COMP_WORDS[COMP_CWORD]}
    COMPREPLY=( $(compgen -W "${targets[@]}" -- $curr_arg ) );
}
complete -F _makefile_targets_01 make
```
