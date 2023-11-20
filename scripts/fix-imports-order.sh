#!/usr/bin/env bash
## Check if gogroup is installed
if ! tool_loc="$(type -p goimports)" || [[ -z ${tool_loc} ]]; then
      echo "goimports is not installed. installing...."
      go install golang.org/x/tools/cmd/goimports@latest
fi

goimports -local=$(go list -m) -w $(find . -type f -name "*.go" | grep -v "vendor/" | grep -v ".git")