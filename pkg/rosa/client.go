// Package rosa provides a way to interact with the Red Hat OpenShift Service on AWS (ROSA) API.
package rosa

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/openshift/rosa/pkg/ocm"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/cluster-api-provider-aws/v2/pkg/cloud/scope"
)

const (
	ocmTokenKey     = "ocmToken"
	ocmAPIURLKey    = "ocmApiUrl"
	ocmClientID     = "ocmClientId"
	ocmClientSecret = "ocmClientSecret"
)

// NewOCMClient creates a new OCM client.
func NewOCMClient(ctx context.Context, rosaScope *scope.ROSAControlPlaneScope) (*ocm.Client, error) {
	sdkConn, err := newOCMRawConnection(ctx, rosaScope)
	if err != nil {
		return nil, err
	}

	result, err := sdkConn.AccountsMgmt().V1().CurrentAccount().Get().Send()

	if err != nil {
		rosaScope.Logger.Error(err, "Cannot esablish OCM client connection, ", result)
		return nil, err
	}

	return ocm.NewClientWithConnection(sdkConn), nil
}

func newOCMRawConnection(ctx context.Context, rosaScope *scope.ROSAControlPlaneScope) (*sdk.Connection, error) {
	logger, err := sdk.NewGoLoggerBuilder().
		Debug(false).
		Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}
	token, url, clientID, clientSecret, err := ocmCredentials(ctx, rosaScope)
	if err != nil {
		return nil, err
	}

	connectionBldr := sdk.NewConnectionBuilder().
		Logger(logger).URL(url)

	if clientID != "" && clientSecret != "" {
		rosaScope.Logger.Info("create connection using client id & secret")
		connectionBldr = connectionBldr.Client(clientID, clientSecret)
	} else if token != "" {
		rosaScope.Logger.Info("create connection using token")
		connectionBldr = connectionBldr.Tokens(token)
	}

	connection, err := connectionBldr.Build()

	if err != nil {
		return nil, fmt.Errorf("failed to create ocm connection: %w", err)
	}

	return connection, nil
}

func ocmCredentials(ctx context.Context, rosaScope *scope.ROSAControlPlaneScope) (string, string, string, string, error) {
	var token, ocmAPIUrl, clientID, clientSecret string

	secret := rosaScope.CredentialsSecret()
	if secret != nil {
		if err := rosaScope.Client.Get(ctx, client.ObjectKeyFromObject(secret), secret); err != nil {
			return "", "", "", "", fmt.Errorf("failed to get credentials secret: %w", err)
		}

		// Keeping the offline token for back compatibility
		// TODO: Should remove the token after offcial deprecation from ocm.
		token = string(secret.Data[ocmTokenKey])
		clientID = string(secret.Data[ocmClientID])
		clientSecret = string(secret.Data[ocmClientSecret])
		ocmAPIUrl = string(secret.Data[ocmAPIURLKey])
	} else {
		// fallback to env variables if secrert is not set
		token = os.Getenv("OCM_TOKEN")
		if ocmAPIUrl = os.Getenv("OCM_API_URL"); ocmAPIUrl == "" {
			ocmAPIUrl = "https://api.openshift.com"
		}
		clientID = os.Getenv(ocmClientID)
		clientSecret = os.Getenv(ocmClientSecret)
	}

	if token == "" && (clientID == "" || clientSecret == "") {
		return "", "", "", "", fmt.Errorf("OCM user credentials not provided, be sure to set the env variable or reference a credentials secret with keys %s , %s", ocmClientID, ocmClientSecret)
	}

	//rosaScope.Logger.Info("ocmCredentials %s, %s , %s, %s", token, ocmAPIUrl, clientID, clientSecret)
	return token, ocmAPIUrl, clientID, clientSecret, nil
}
