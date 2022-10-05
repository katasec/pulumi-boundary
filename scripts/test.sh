#!/bin/bash

tagging_stuff()  {
    export VERSION="v0.0.1"
    git tag -d "${VERSION}"
    git push --delete origin "${VERSION}"

    git tag -a "v0.0.1" -m "v0.0.1"
    git push origin tags "${VERSION}"
}


# Build pulumi resource binary
REPO_ROOT=`git rev-parse --show-toplevel`
cd "$REPO_ROOT/provider/cmd/pulumi-resource-boundary"
go build -o "$REPO_ROOT/bin" "$REPO_ROOT/provider/cmd/pulumi-resource-boundary"
export VERSION=`git tag --points-at HEAD`

# Create folder sctructure for release
folder_name="pulumi-resource-boundary-${VERSION}-linux-amd64"
export tmp=`mktemp -d`
mkdir -p "$tmp/$folder_name"
cp "$REPO_ROOT/bin/pulumi-resource-boundary" "$tmp/$folder_name"
cd $tmp
tar -czvf "$folder_name.tar.gz" $folder_name

tar -tzvf "$folder_name.tar.gz" $folder_name

cd $REPO_ROOT
gh release create "$VERSION" "$tmp/$folder_name.tar.gz" -t "$VERSION"
