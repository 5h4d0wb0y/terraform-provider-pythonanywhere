#!/bin/bash

BUILD_DIR="build"
DATE=$(date +%F)
CURRENT_VERSION=$(git describe --tags --always --dirty --match=v* 2> /dev/null || cat .version 2> /dev/null || echo v0)
REPO=${REPO:-$(pwd | cut -d '/' -f 6)}
GHUSER=${GHUSER:-$(git config --get user.name)}
LAST_TAG=$(git describe --abbrev=0 --tags)
ARCHS=${ARCHS:-"linux/amd64 linux/arm darwin/amd64 windows/amd64"}
DRAFT=${DRAFT:-true}

step "[+] Retrieving dependencies ..."
go get github.com/aktau/github-release
go get github.com/mitchellh/gox

echo "[+] Cleaning up build directory if present ..."
rm -rf ${BUILD_DIR}

echo -n "[?] Current version is $CURRENT_VERSION, select new version: "
read NEW_VERSION
echo "[+] Creating version $NEW_VERSION ...\n"

echo "[+] Fetching remote tags ..."
NEW=()
FIXES=()
MISC=()
ALL=()
OUTPUT=$(git log $LAST_TAG..HEAD --oneline)
IFS=$'\n' LINES=($OUTPUT)
for LINE in "${LINES[@]}"; do
    LINE=$(echo "$LINE" | sed -E "s/^[[:xdigit:]]+\s+//")
    if [[ $LINE = *"new:"* ]]; then
        LINE=$(echo "$LINE" | sed -E "s/^new: //")
        NEW+=("$LINE")
    elif [[ $LINE = *"fix:"* ]]; then
        LINE=$(echo "$LINE" | sed -E "s/^fix: //")
        FIXES+=("$LINE") 
    elif [[ $LINE != *"i did not bother commenting"* ]] && [[ $LINE != *"Merge "* ]]; then 
        echo "MISC LINE =$LINE"
        LINE=$(echo "$LINE" | sed -E "s/^[a-z]+: //")
        MISC+=("$LINE")
    fi
    ALL+=("$LINE")
done

echo "[+] Creating release text description ..."
echo "" > release.tmp
if [ -n "$NEW" ]; then
    echo -e "**New Features**\n" >> release.tmp
    for l in "${NEW[@]}"
    do
        echo "* $l" >> release.tmp
    done
fi
echo "" >> release.tmp
if [ -n "$FIXES" ]; then
    echo -e "\n**Fixes**\n" >> release.tmp
    for l in "${FIXES[@]}"
    do
        echo "* $l" >> release.tmp
    done
fi
echo "" >> release.tmp
if [ -n "$MISC" ]; then
    echo -e "\n**Misc**\n" >> release.tmp
    for l in "${MISC[@]}"
    do
        echo "* $l" >> release.tmp
    done
fi

echo "[+] Updating changelog ..."
echo -e "\n$(cat CHANGELOG.md)" > CHANGELOG.md
for l in "${ALL[@]}"
do
    echo -e "   * $l\n$(cat CHANGELOG.md)" > CHANGELOG.md
done
echo -e "# $CURRENT_VERSION/ $DATE\n\n$(cat CHANGELOG.md)" > CHANGELOG.md

echo "[+] Releasing v$NEW_VERSION ..."
git commit -m "Releasing v$NEW_VERSION"
git push
git tag -a v$NEW_VERSION -m "Release v$NEW_VERSION"
git push origin v$NEW_VERSION

echo "[+] Compiling program ..."
mkdir ${BUILD_DIR}
gox -ldflags="-X main.version=${VERSION}" -osarch="${ARCHS}" \
    -output="${BUILD_DIR}/{{.Dir}}_{{.OS}}_{{.Arch}}"

echo "[+] Generating binary SHASUMs ..."
cd ${BUILD_DIR}
sha256sum * > SHA256SUMS

echo "[+] Packing archives ..."
for file in *; do
	if [ "${file}" = "SHA256SUMS" ]; then
		continue
	fi

	if [[ ${file} == *linux* ]]; then
		tar -czf "${file%%.*}.tar.gz" "${file}"
	else
		zip "${file%%.*}.zip" "${file}"
	fi

	rm "${file}"
done

if [ -z "${GITHUB_TOKEN}" ]; then
	echo '[x] Please set $GITHUB_TOKEN environment variable to avoid having to enter it for each new release!'
	echo -n "[?] Put the GitHub Token: "
    read GITHUBTOKEN
    export GITHUB_TOKEN="$GITHUBTOKEN"
fi

if [[ ${DRAFT} == "true" ]]; then
	echo "[+] Creating a drafted release ..."
	github-release release --user ${GHUSER} \
                           --repo ${REPO} \
                           --tag "v$NEW_VERSION" \
                           --name "v$NEW_VERSION" \
                           --draft || true
else
	echo "[+] Creating a published release ..."
	github-release release --user ${GHUSER} \
                           --repo ${REPO} \
                           --tag "v$NEW_VERSION" \
                           --name "v$NEW_VERSION" \
                           --description ${cat release.tmp} || true
fi

echo "[+] Uploading build assets ..."
for file in *; do
	echo "- ${file}"
	github-release upload --user ${GHUSER} \
                          --repo ${REPO} \
                          --tag "v$NEW_VERSION" \
                          --name ${file} \
                          --file ${file}
done

echo "[+] Cleaning up build directory ..."
rm -rf ${BUILD_DIR} release.tmp

echo
echo "[+] All done, v$NEW_VERSION released!"
