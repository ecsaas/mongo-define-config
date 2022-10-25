package mongodb

import (
	"context"
	"fmt"
	"github.com/ecsaas/mongo-define-config/DEFINE_VARIABLES/MONGO_ENV"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Mdb *mongo.Client
var MdbCancelFunc func()
var MdbErr error

func MongoDB() {
	var applyURI = fmt.Sprintf(
		"%s://%s:%s@%s/admin",
		os.Getenv(MONGO_ENV.MDB_SRV),
		url.QueryEscape(os.Getenv(MONGO_ENV.MDB_USER)),
		url.QueryEscape(os.Getenv(MONGO_ENV.MDB_PASS)),
		os.Getenv(MONGO_ENV.MDB_CONN),
	)
	Mdb, MdbErr = mongo.NewClient(options.Client().ApplyURI(applyURI))
	if MdbErr != nil {
		return
	}
	var ctx context.Context
	ctx, MdbCancelFunc = context.WithTimeout(context.Background(), 10*time.Second)
	MdbErr = Mdb.Connect(ctx)
}
