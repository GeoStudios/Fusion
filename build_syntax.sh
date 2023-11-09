#!/bin/bash

# Change to the 'syntax-highlighter' directory
cd syntax-highlighter

# Run 'npx vsce package'
npx vsce package

# Change back to the parent directory
# shellcheck disable=SC2103
cd ..
