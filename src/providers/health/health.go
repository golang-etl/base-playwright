package health

import (
	"context"
	"net/http"
	"time"

	packagegeneralinterfaces "github.com/golang-etl/package-general/src/interfaces"
	packagehttpconsts "github.com/golang-etl/package-http/src/consts"
	packagehttpinterfaces "github.com/golang-etl/package-http/src/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type HealthProvider struct {
	CfgGoModuleName string
	CfgDebug        bool
	MongoClient     *mongo.Client
}

func (provider HealthProvider) GetHealth(shared *packagegeneralinterfaces.Shared) packagehttpinterfaces.Response {
	return packagehttpinterfaces.Response{
		StatusCode: http.StatusOK,
		Headers:    packagehttpconsts.HeaderContentType.JSON,
		Body: map[string]interface{}{
			"Status": "OK",
			"Module": provider.CfgGoModuleName,
			"Database": map[string]interface{}{
				"MongoDB": map[string]interface{}{
					"Ping": provider.GetMongoDBPing(),
				},
			},
		},
	}
}

func (provider HealthProvider) GetMongoDBPing() *int64 {
	start := time.Now()
	err := provider.MongoClient.Ping(context.TODO(), nil)

	if err != nil {
		return nil
	}

	elapsed := time.Since(start).Milliseconds()

	return &elapsed
}
