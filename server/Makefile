
# Assumptions: 
# The directory of this file is under aks-rp. This is to find aksbuilder.sh.
# The aks-rp directory is under $HOME. The azureconfig.yaml is at aks-rp/bin/azureconfig.yaml.
# This is to find azureconfig.yaml used by deploy.
# Modify the path in this file if the assumption of this file is broken.
bootstrap:
	../../hack/aksbuilder.sh bootstrap -w ./

tidy:
	../../hack/aksbuilder.sh tidy -w ./

build:
	../../hack/aksbuilder.sh build -w ./

deploy:
	../../hack/aksbuilder.sh localdeploy --workspace ./ --usespoperator --e2eazureconfig ~/aks-rp/bin/azureconfig.yaml
