# Completion for aligo
# This completion is automatically generated

_aligo() {
  local cur prev cmds opts show_files

  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"

  cmds="check view"
  opts="--arch --struct --tags --pager --no-color --help --version"
  show_files="true"
  file_glob="{{FILE_GLOB}}"



  if [[ $cur == -* ]] ; then
    COMPREPLY=($(compgen -W "$opts" -- "$cur"))
    return 0
  fi

  _aligo_filter "$cmds" "$opts" "$show_files" "$file_glob"
}

_aligo_filter() {
  local cmds="$1"
  local opts="$2"
  local show_files="$3"
  local file_glob="$4"

  local cmd1 cmd2

  for cmd1 in $cmds ; do
    for cmd2 in ${COMP_WORDS[*]} ; do
      if [[ "$cmd1" == "$cmd2" ]] ; then
        if [[ -z "$show_files" ]] ; then
          COMPREPLY=($(compgen -W "$opts" -- "$cur"))
        else
          _filedir "$file_glob"
        fi

        return 0
      fi
    done
  done

  if [[ -z "$show_files" ]] ; then
    COMPREPLY=($(compgen -W "$cmds" -- "$cur"))
    return 0
  fi

  _filedir "$file_glob"
}

complete -F _aligo aligo -o filenames
