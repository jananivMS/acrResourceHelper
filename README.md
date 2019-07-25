# acrResourceHelper
Helper functions to create delete ACR using Go SDK


1. Create a setenv.sh file in that folder with the below information
`

export AZURE_CLIENT_ID=<Client ID of your service principal>

export AZURE_CLIENT_SECRET=<Client Secret of your service principal>

export AZURE_TENANT_ID=<Azure tenant ID where you want to deploy the ACR>

export AZURE_SUBSCRIPTION_ID=<Azure Subscription ID where you want to deploy the ACR>`

The client ID and client secret is of the service principal in Azure that you have created and have given write access to your resource group where you want to create the ACR. You will need to do this in the Azure portal before running this.

2. Run the above file using `source setenv.sh` in the command prompt where you would be executing the project from. This sets these values as environment variables for the program to read from.
