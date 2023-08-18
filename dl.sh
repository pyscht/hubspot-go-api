#!/bin/bash

function docker_run() {
  docker run --rm \
    -v "${PWD}":/app \
    -w /app \
    "${@}"
}

function generate_openapi() {
  local src=''
  local out=''
  local subdir=''
  local jopts='-Xss4M'

  local args=''

  while (($# > 0)); do
    local arg="${1}"; shift

    case "${arg}" in
      -s | --source) src="${1}"; shift ;;
      -o | --output) out="${1}"; shift ;;
      -d | --subdir) subdir="${1}"; shift ;;
      -j | --java-opts) jopts="${1}"; shift ;;
      *) args="${args} ${arg#'-X'}"
    esac
  done

  if [[ -z "${out}" ]]; then
    echo "ERROR: -o | --out is required"
    exit 1
  fi

  local root
  local tmpdir
  root="$(git rev-parse --show-toplevel)"
  tmpdir="$(mktemp -d)"

  # shellcheck disable=SC2115
  rm -rf "${root}/${out}/${subdir}"
  mkdir -p "${root}/${out}"

  docker_run \
    -v "${tmpdir}:/output" \
    -e JAVA_OPTS="${jopts}" \
    openapitools/openapi-generator-cli:v6.6.0 \
    generate \
      -i "${src}" \
      -o '/output' \
      ${args#' '}

  cp -R "${tmpdir}/${subdir}" "${root}/${out}/"

  rm -rf "${tmpdir}"
}

function generate_api_client_go {
  local service
  service=$1

  generate_openapi -g go \
     -s "${service}/openapi-spec.json" \
     -o "${service}/generated/" \
     -c "${service}/config.go-client.yaml" \
     --skip-validate-spec

  mv "${service}/generated/"* "${service}/" && \
    mv "${service}/generated/".* "${service}/" 2>/dev/null

  rm -fr "${service}/generated"
}

# Download JSON
JSON_URL="https://api.hubspot.com/api-catalog-public/v1/apis"
TEMP_JSON="temp.json"
curl -o "$TEMP_JSON" "$JSON_URL"

rm -fr ./apis/*

jq -c -r '.results[] ' "$TEMP_JSON" | while IFS=$'\t' read -r api; do

  name=$(echo "$api" | jq -r ".name")

  echo "$api" | jq -c -r '.features | to_entries[] | "\(.key)\t\(.value.openAPI)"' | while IFS=$'\t' read -r feature url; do
    module=$(echo "$name" | tr ' ' '-' | tr '[:upper:]' '[:lower:]')
    package=$(echo "$feature" | tr ' ' '_' | tr '-' '_' | tr '[:upper:]' '[:lower:]')
    api_path="${module}/${package}"
    echo "api: $api_path"
    echo "url: $url"
    mkdir -p "apis/$api_path"
    curl -o "apis/$api_path/openapi-spec.json" "$url"
    cp config.go-client.yaml "apis/${api_path}/"
    echo -e "gitUserId: pyscht\ngitRepoId: hubspot-go-api/apis/${module}\npackageName: ${package}" >> "apis/${api_path}/config.go-client.yaml"
    generate_api_client_go "apis/$api_path"
    (cd "apis/$api_path" && go mod download && go fmt)
    go work use "apis/$api_path"
  done
done

rm "$TEMP_JSON"
