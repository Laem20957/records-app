#!/bin/sh

cp ./.scripts/pre-commit-hook ./.git/hooks/pre-commit
chmod +x ./.git/hooks/pre-commit
cp ./.scripts/post-commit-hook ./.git/hooks/post-commit
chmod +x ./.git/hooks/post-commit
