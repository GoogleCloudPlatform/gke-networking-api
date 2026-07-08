#!/usr/bin/env bash

# Copyright 2026 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

# Go to the root of the repository
REPO_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
cd "${REPO_ROOT}"

# Print help message
function print_usage() {
  echo "Usage: $0 [options]"
  echo "Options:"
  echo "  -y, --yes          Non-interactive mode: auto-approve tag creation and push"
  echo "  -d, --dry-run      Dry run mode: calculate and print the next tag, but do not create it"
  echo "  -h, --help         Show this help message"
}

YES_MODE=false
DRY_RUN=false

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case "$1" in
    -y|--yes)
      YES_MODE=true
      shift
      ;;
    -d|--dry-run)
      DRY_RUN=true
      shift
      ;;
    -h|--help)
      print_usage
      exit 0
      ;;
    *)
      echo "Unknown option: $1"
      print_usage
      exit 1
      ;;
  esac
done

# 1. Determine remote
if git remote | grep -q '^upstream$'; then
  REMOTE="upstream"
else
  REMOTE="origin"
fi

echo "Using remote: $REMOTE"

# 2. Fetch latest tags and commits
echo "Fetching tags and updates from $REMOTE..."
git fetch "$REMOTE" --tags || echo "Warning: failed to fetch tags/updates from $REMOTE. Proceeding with local cache."

# 3. Get the newest commit on the main branch of the remote
TARGET_BRANCH="main"
COMMIT=$(git rev-parse "refs/remotes/$REMOTE/$TARGET_BRANCH" 2>/dev/null || git rev-parse "$TARGET_BRANCH" 2>/dev/null || true)

if [[ -z "$COMMIT" ]]; then
  echo "Error: Could not find target branch '$TARGET_BRANCH' or 'refs/remotes/$REMOTE/$TARGET_BRANCH'"
  exit 1
fi

echo "Target commit to tag on $REMOTE/$TARGET_BRANCH: $COMMIT"
git log -1 --format="  Author:  %an <%ae>%n  Date:    %cd%n  Subject: %s" "$COMMIT"
echo

# 4. Find the latest tag merged into this commit
LATEST_TAG=$(git tag -l "v*" --merged "$COMMIT" | sort -V | tail -n 1)

NEXT_TAG=""

if [[ -z "$LATEST_TAG" ]]; then
  echo "Error: No existing tags found merged into commit $COMMIT."
  exit 1
else
  echo "Latest tag merged into this commit: $LATEST_TAG"
  
  # Strip 'v' prefix
  LATEST_VERSION="${LATEST_TAG#v}"
  
  # Split version into components
  IFS='.' read -r major minor patch <<< "$LATEST_VERSION"
  
  # Verify version format is numeric
  if [[ "$major" =~ ^[0-9]+$ ]] && [[ "$minor" =~ ^[0-9]+$ ]] && [[ "$patch" =~ ^[0-9]+$ ]]; then
    NEXT_PATCH=$((patch + 1))
    NEXT_TAG="v${major}.${minor}.${NEXT_PATCH}"
    echo "Incrementing patch version from $LATEST_TAG to $NEXT_TAG."
  else
    echo "Error: Latest tag format is not standard semver: $LATEST_TAG."
    exit 1
  fi
fi

if [ "$DRY_RUN" = true ]; then
  echo "Dry run mode active: calculated tag would be $NEXT_TAG."
  exit 0
fi

# 5. Check if the NEXT_TAG already exists in git
if git rev-parse "$NEXT_TAG" >/dev/null 2>&1; then
  echo "Error: Tag $NEXT_TAG already exists!"
  exit 1
fi

# 6. Ask for confirmation unless YES_MODE is active
if [ "$YES_MODE" = false ]; then
  read -p "Do you want to tag commit $COMMIT with $NEXT_TAG? (y/N): " -r CONFIRM
  if [[ ! "$CONFIRM" =~ ^[Yy]$ ]]; then
    echo "Aborted tagging."
    exit 0
  fi
fi

# 7. Create tag
echo "Tagging commit $COMMIT with $NEXT_TAG..."
git tag "$NEXT_TAG" "$COMMIT"

# 8. Ask for confirmation to push unless YES_MODE is active
if [ "$YES_MODE" = false ]; then
  read -p "Do you want to push the tag $NEXT_TAG to $REMOTE? (y/N): " -r PUSH_CONFIRM
  if [[ "$PUSH_CONFIRM" =~ ^[Yy]$ ]]; then
    echo "Pushing tag $NEXT_TAG to $REMOTE..."
    git push "$REMOTE" "$NEXT_TAG"
  else
    echo "Tag created locally. You can push it manually using: git push $REMOTE $NEXT_TAG"
  fi
else
  echo "Pushing tag $NEXT_TAG to $REMOTE..."
  git push "$REMOTE" "$NEXT_TAG"
fi
