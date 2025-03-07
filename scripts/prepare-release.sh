#!/usr/bin/env bash
set -e
set -o pipefail

if [[ -z "$OLM_TYPE" ]]; then
  OLM_TYPE="integreatly-operator"
fi

case $OLM_TYPE in
  "integreatly-operator")
    PREVIOUS_VERSION=$(grep $OLM_TYPE packagemanifests/$OLM_TYPE/$OLM_TYPE.package.yaml | awk -F v '{print $2}') || echo "No previous version"
    OPERATOR_TYPE=rhmi
    ;;
  "managed-api-service")
    PREVIOUS_VERSION=$(grep $OLM_TYPE packagemanifests/$OLM_TYPE/$OLM_TYPE.package.yaml | awk -F v '{print $3}') || echo "No previous version"
    OPERATOR_TYPE=rhoam
    ;;
  *)
    echo "Invalid OLM_TYPE set"
    echo "Use \"integreatly-operator\" or \"managed-api-service\""
    exit 1
    ;;
esac

if [[ -z "$ORG" ]]; then
  ORG="integreatly"
else
  ORG="$ORG"
fi

# Optional environment variable to set a different Kustomize path. If this
# variable is not set, it will use the one from the $PATH or install Kustomize
if [[ -z $KUSTOMIZE_PATH ]]; then
  KUSTOMIZE="/usr/local/bin/kustomize"
else
  KUSTOMIZE="/usr/local/bin/kustomize"
fi

# Path to gofmt
if [[ -z $GOROOT ]]; then
  GOFMT="/usr/local/go/bin/gofmt"
else
  GOFMT="$GOROOT/bin/gofmt"
fi

create_new_csv() {

  if [[ -z "$PREVIOUS_VERSION" ]]
    then
      "${KUSTOMIZE[@]}" build config/manifests-$OPERATOR_TYPE | operator-sdk generate packagemanifests --kustomize-dir=config/manifests-$OPERATOR_TYPE --output-dir packagemanifests/$OLM_TYPE --version $VERSION --default-channel --channel rhmi 
    else
      "${KUSTOMIZE[@]}" build config/manifests-$OPERATOR_TYPE | operator-sdk generate packagemanifests --kustomize-dir=config/manifests-$OPERATOR_TYPE --output-dir packagemanifests/$OLM_TYPE --version $VERSION --default-channel --channel rhmi --from-version $PREVIOUS_VERSION
  fi
}

update_csv() {
  "${KUSTOMIZE[@]}" build config/manifests-$OPERATOR_TYPE | operator-sdk generate packagemanifests --kustomize-dir=config/manifests-$OPERATOR_TYPE --output-dir packagemanifests/$OLM_TYPE --version $VERSION --default-channel --channel rhmi 
}

# The base CSV is used to generate the final CSV by combining it with the other operator
# manifests. In operator-sdk v1.2.0, the replaces field of the new CSV is set from
# the current version of **the base CSV**, so we need to update the base CSV in order
# for the replaces field to be set when generating the next release
update_base_csv() {
  yq w -i config/manifests-$OPERATOR_TYPE/bases/$OLM_TYPE.clusterserviceversion.yaml metadata.name $OLM_TYPE.v$VERSION
  yq w -i config/manifests-$OPERATOR_TYPE/bases/$OLM_TYPE.clusterserviceversion.yaml spec.version $VERSION
}

set_version() {
  if [[ -z "$PREVIOUS_VERSION" ]]
    then
      echo "No previous version please set correct values in the Makefile and version/version.go files"
    else
      case $OLM_TYPE in
        "integreatly-operator")
          "${SED_INLINE[@]}" -E "s/RHMI_TAG\s+\?=\s+$PREVIOUS_VERSION/RHMI_TAG \?= $VERSION/g" Makefile
          "${SED_INLINE[@]}" -E "s/version\s+=\s+\"$PREVIOUS_VERSION\"/version = \"$VERSION\"/g" version/version.go
          ;;
        "managed-api-service")
          "${SED_INLINE[@]}" -E "s/RHOAM_TAG\s+\?=\s+$PREVIOUS_VERSION/RHOAM_TAG \?= $VERSION/g" Makefile
          "${SED_INLINE[@]}" -E "s/managedAPIVersion\s+=\s+\"$PREVIOUS_VERSION\"/managedAPIVersion = \"$VERSION\"/g" version/version.go
          ;;
        *)
          echo "No version found for install type : $(OLM_TYPE)"
          ;;
      esac
  fi
}

set_installation_type() {
  if [[ -z "$PREVIOUS_VERSION" ]]
    then
      echo "No previous version please set correct values in the Makefile and version/version.go files"
    else
      case $OLM_TYPE in
        "integreatly-operator")
          echo "using default INSTALLATION_TYPE found in deploy/operator.yaml"
          ;;
        "managed-api-service")
          yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" --tag '!!str' spec.install.spec.deployments[0].spec.template.spec.containers[0].env.'(name==INSTALLATION_TYPE)'.value managed-api
          ;;
        *)
          echo "No INSTALLATION_TYPE found for install type : $(OLM_TYPE)"
          echo "using default INSTALLATION_TYPE found in deploy/operator.yaml"
          ;;
      esac
  fi
}

set_descriptions() {
  case $OLM_TYPE in
   "integreatly-operator")
      echo "using default descriptions"
      ;;
    "managed-api-service")
      echo "Updating descriptions"
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/integreatly.org_rhmis.yaml" --tag '!!str' spec.validation.openAPIV3Schema.description 'RHOAM is the Schema for the RHOAM API'
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/integreatly.org_rhmis.yaml" --tag '!!str' spec.validation.openAPIV3Schema.properties.spec.description 'RHOAMSpec defines the desired state of Installation'
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/integreatly.org_rhmis.yaml" --tag '!!str' spec.validation.openAPIV3Schema.properties.status.description 'RHOAMStatus defines the observed state of Installation'
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/integreatly.org_rhmiconfigs.yaml" --tag '!!str' spec.validation.openAPIV3Schema.description 'RHOAMConfig is the Schema for the rhoamconfigs API'
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/integreatly.org_rhmiconfigs.yaml" --tag '!!str' spec.validation.openAPIV3Schema.properties.spec.description 'RHOAMConfigSpec defines the desired state of RHOAMConfig'
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/integreatly.org_rhmiconfigs.yaml" --tag '!!str' spec.validation.openAPIV3Schema.properties.status.description 'RHOAMConfigStatus defines the observed state of RHOAMConfig'
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/integreatly.org_rhmiconfigs.yaml" --tag '!!str' spec.validation.openAPIV3Schema.properties.status.properties.upgradeAvailable.properties.targetVersion.description 'target-version: string, version of incoming RHOAM Operator'

      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" --tag '!!str' spec.customresourcedefinitions.owned[1].description 'RHOAM is the Schema for the RHOAM API'
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" --tag '!!str' spec.customresourcedefinitions.owned[1].displayName 'RHOAM installation'

      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" --tag '!!str' spec.customresourcedefinitions.owned[0].description 'RHOAMConfig is the Schema for the rhoamconfigs API'
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" --tag '!!str' spec.customresourcedefinitions.owned[0].displayName 'RHOAMConfig'

      ;;
  esac
}

set_clusterPermissions() {
  case $OLM_TYPE in
   "integreatly-operator")
      echo "using default permissions"
      ;;
    "managed-api-service")
      echo "Updating permissions"
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" --tag '!!str' spec.maintainers[0].email 'rhoam-support@redhat.com'
      yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" --tag '!!str' spec.maintainers[0].name 'rhoam'
      ;;
  esac
}

set_images() {
  case $OLM_TYPE in
   "integreatly-operator")
  : "${IMAGE_TAG:=v${SEMVER}}"
  yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" spec.install.spec.deployments.[0].spec.template.spec.containers[0].image quay.io/$ORG/$OLM_TYPE:$IMAGE_TAG
  yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" metadata.annotations.containerImage quay.io/$ORG/$OLM_TYPE:$IMAGE_TAG
  ;;
  "managed-api-service")
   : "${IMAGE_TAG:=rhoam-v${SEMVER}}"
  yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" spec.install.spec.deployments.[0].spec.template.spec.containers[0].image quay.io/$ORG/$OLM_TYPE:$IMAGE_TAG
  yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" metadata.annotations.containerImage quay.io/$ORG/$OLM_TYPE:$IMAGE_TAG

  ;;
  esac
}

set_csv_service_affecting_field() {
  local value=$1
  echo "Update CSV for release $SEMVER to be 'serviceAffecting: $value'"
  yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" --tag '!!str' metadata.annotations.serviceAffecting "$value"
}

# Due to a quirk in operator-sdk v1.2.0, the `replaces` field in the generated CSV is taken from the current version in
# the base CSV (config/manifests-(rhoam|rhmi)/bases). When generating the CSV from a new version (e.g. from 1.5.0
# to 1.6.0), this script updates the version in the base CSV to the new version, so the next time the script runs it
# uses it as the `replaces` field value. However, when cutting an RC from the same version (1.5.0-rc1 to 1.5.0-rc2), the
# version in the base CSV is the same as the new version, so operator-sdk generates a CSV without the `replaces` field.
# This function addresses that by adding the `replaces` field if it was removed.
check_csv_replaces_field() {
  if [[ -z $(yq r "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" spec.replaces) ]]; then
    yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" spec.replaces "${PREVIOUS_REPLACES_VALUE}"
  fi
}

# Sets the related images in the CSV for RHOAM
set_related_images() {
  echo "Adding related images to the CSV"
  containerImageField="""[
  """
  position=0
  length=$(yq r -j ./products/products.yaml | jq -r '.products' | jq length)
  # Get supported components
  for (( i=0; i<${length}; i++))
  do
    product_dir=$(yq r ./products/products.yaml "products[$i].manifestsDir")
    if [[ $(yq r ./products/products.yaml "products[$i].installType") == *"rhoam"* && $(yq r ./products/products.yaml "products[$i].quayScan") == true ]]; then
      # Read component version
      component_version=$(grep currentCSV manifests/$product_dir/*.package.yaml | awk -F v '{print $2}')

      # Read component name
      component_name=$(yq r -j ./manifests/$product_dir/${component_version}/*.clusterserviceversion.yaml | jq '.metadata.name' |  tr -d '"')
      
      # Read image from the component version but only select quay.io or redhat.registry
      component_image=$(yq r -j ./manifests/$product_dir/${component_version}/*.clusterserviceversion.yaml | jq '.spec.install.spec.deployments[0].spec.template.spec.containers' | jq '.[] | select((.image|test("quay.")) or (.image|test("registry.redhat"))) | .image' |  tr -d '"')

      containerImageField="$containerImageField{\"component_name\":\"${component_name}\",\"component_url\":\"${component_image}\"},"
      position=$((position+1))

      # Check if the CSV of the component has the relatedImages set, if it does, populate RHOAM CSV with it.
      relatedImagesLength=$(yq r -j ./manifests/$product_dir/${component_version}/*.clusterserviceversion.yaml | jq -r '.spec.relatedImages' | jq length)

      # Adding generic related images
      if [[ $relatedImagesLength != 0 ]]; then
        for (( y=0; y<$relatedImagesLength; y++))
        do
          relatedImageName=$(yq r -j ./manifests/$product_dir/${component_version}/*.clusterserviceversion.yaml | jq -r ".spec.relatedImages[$y].name")
          relatedImageURL=$(yq r -j ./manifests/$product_dir/${component_version}/*.clusterserviceversion.yaml | jq -r ".spec.relatedImages[$y].image")
          if [[ "$relatedImageName" != "redis-32-rhel7" ]] && [[ "$relatedImageName" != "mysql-57-rhel7" ]] && [[ "$relatedImageName" != "openshift-cli" ]]; then
            containerImageField="$containerImageField{\"component_name\":\"${relatedImageName}\",\"component_url\":\"${relatedImageURL}\"},"
            position=$((position+1))
          fi
        done
      fi

      # Adding KC related image
      if [[ "$component_name" == *"keycloak-operator"* ]]; then
        kcRelatedImageName=$(yq r -j ./manifests/$product_dir/${component_version}/*.clusterserviceversion.yaml | jq '.spec.install.spec.deployments[0].spec.template.spec.containers[0].env[0].name' |  tr -d '"')
        kcRelatedImageURL=$(yq r -j ./manifests/$product_dir/${component_version}/*.clusterserviceversion.yaml | jq '.spec.install.spec.deployments[0].spec.template.spec.containers[0].env[0].value' |  tr -d '"')
        containerImageField="$containerImageField{\"component_name\":\"${kcRelatedImageName}\",\"component_url\":\"${kcRelatedImageURL}\"},"
      fi
    fi
  done
  containerImageRemovedLastCharacter=$(echo "${containerImageField::-1}")
  containerImageField="$containerImageRemovedLastCharacter]"
  yq w -i "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" --tag '!!str' metadata.annotations.containerImages "$containerImageField"
}

if [[ -z "$SEMVER" ]]; then
 echo "ERROR: no SEMVER value set"
 exit 1
fi

if [[ $SEMVER =~ ^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(-(0|[1-9][0-9]*|[0-9]*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9][0-9]*|[0-9]*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\+[0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*)?$ ]]; then
  echo "Valid version string: ${SEMVER}"
else
  echo "Error: Invalid version string: ${SEMVER}"
  exit 1
fi

VERSION=$(echo "$SEMVER" | awk -F - '{print $1}')

# Set sed -i as it's different for mac vs gnu
if [[ $(uname) = Darwin ]]; then
  SED_INLINE=(sed -i '')
else
  SED_INLINE=(sed -i)
fi

# The `projectName` field in the PROJECT file is used by the operator-sdk CLI
# to generate the CSV. In order to be compatible with both types of CSVs
# (RHMI and RHOAM), we need to temporarily set the `projectName` to the desired
# OLM type, and save the current value in order to reset it when we're done
current_project_name=$(yq r PROJECT projectName)
yq w -i PROJECT projectName $OLM_TYPE

# We have a new version so generate the csv
if [[ "$VERSION" != "$PREVIOUS_VERSION" ]]; then
  create_new_csv
  set_version
else
  # Save the previous value of the `replaces` field. Needed to repopulate the `replaces` field later if removed.
  PREVIOUS_REPLACES_VALUE=$(yq r "packagemanifests/$OLM_TYPE/${VERSION}/$OLM_TYPE.clusterserviceversion.yaml" spec.replaces)
  update_csv
fi

set_installation_type
set_descriptions
set_clusterPermissions
set_images

if [[ -n "$SERVICE_AFFECTING" ]]; then
 set_csv_service_affecting_field "$SERVICE_AFFECTING"
fi

if [[ "${OLM_TYPE}" == "managed-api-service" ]]; then
  set_related_images
fi

update_base_csv
check_csv_replaces_field

# Reset the project name
yq w -i PROJECT projectName $current_project_name

# Ensure the RHMI package is `integreatly`: The operator-sdk CLI will take the
# package name from the PROJECT file, so in the case of RHMI it will set it
# incorrectly to `integreatly-operator`
yq w -i packagemanifests/integreatly-operator/integreatly-operator.package.yaml packageName integreatly

# Ensure code is formatted correctly
"${GOFMT[@]}" -w `find . -type f -name '*.go' -not -path "./vendor/*"`