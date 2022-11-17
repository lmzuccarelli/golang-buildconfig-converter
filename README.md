# A simple openshift buildconfig converter to shipwright build manifests

## Intro

This is a simple golang projects that takes a set of openshift buildconfigs and converts them to the equivalent build manifests used in 'shipwright'.

This was succesfully tested with golang 1.19.3. 
I have tried to keep coverage over 90%


### Clone the repository and build

```bash
git clone git@github.com:lmzuccarelli/golang-buildconfig-converter

cd golang-buildconfig-converter-
make clean
make build

```

## Usage

Generate the relevant build manifests by pointing to a directory with buildconfigs

Execute the following command


```bash
# setup a simple config.yaml file

cat << EOF > config.yaml
apiVersion: 0.0.1
kind: Config
metadata:
  name: converter-config
spec:
  workingDirectory: working-dir
  buildConfigPath: ../okd-payload-pipeline/buildconfigs
EOF


# this will generate the shipwright builds to the working directory
./build/convert -c config.yaml

```
