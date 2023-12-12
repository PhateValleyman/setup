_setup_completion()
{
    local cur opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"

    # Define available options
    opts="help alias defaults functions paths profile"

    # Generate completions
    COMPREPLY=($(compgen -W "${opts}" -- "${cur}"))
}

# Register the completion function for the 'setup' command
complete -F _setup_completion setup
