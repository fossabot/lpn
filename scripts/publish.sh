#!/bin/bash

function git_branch_name() {
  echo $(git symbolic-ref --short HEAD)
}

readonly BRANCH="${TRAVIS_BRANCH:-$(git_branch_name)}"

function publish_website() {
  local devEnvironment="dev"
  local remote="liferay.io"

  cd wedeploy
  we remote set liferay.io liferay.io
  we remote default liferay.io
  echo "$WE_TOKEN" | we login -r "${remote}"

  if [[ "$BRANCH" == "master" ]]; then
    we deploy -r "${remote}" -p lpn
  elif [[ "$BRANCH" == "develop" ]]; then
    echo "INFO:
    Deploying from develop branch will push the website to the Dev environment for this project
    "
    we deploy -r "${remote}" -p lpn -e "${devEnvironment}"
  else
    echo "We cannot deploy the website from a branch different than master or develop"
    exit 1
  fi
}

publish_website