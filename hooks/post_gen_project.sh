#! /usr/bin/env bash

# Get Current Time from Google
curl -I 'https://google.com/' 2>/dev/null | grep -i '^date:' | sed 's/^[Dd]ate: //g'

# You can replace this with any kind of script to reachout post generation
# For example you could trigger a terraform cloud run
# curl \
#   --request POST \
#   --header "Authorization: Bearer XXX_TOKEN_XXX" \
#   --header "Content-Type: application/vnd.api+json" \
#   --data @payload.json \
#   https://app.terraform.io/api/v2/workspaces/ws-XXX_WORKSPACE_ID_XXX/run-triggers