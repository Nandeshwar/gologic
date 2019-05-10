

"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

var db dynamodbiface.DynamoDBAPI = dynamodb.New(sess)
	repo, err := repo.New(db, tableName, int64(recordFetchLimit))



//go:generate mockgen -destination mock/mock_repo/mock_repo.go -source repo.go DataRepository
type DataRepository interface {
	ReadLocation(deviceID string, daysInPast int) ([]Location, error)
	PutArbitraryJSON(arbitraryJSON map[string]interface{}) (*dynamodb.PutItemOutput, error)
}

type DynamoDB struct {
	db    dynamodbiface.DynamoDBAPI
	table string
	Limit int64
}

func New(db dynamodbiface.DynamoDBAPI, table string, limit int64) (*DynamoDB, error) {
	return &DynamoDB{
		db:    db,
		table: table,
		Limit: limit,
	}, nil
}





func (r *DynamoDB) queryBySortKeyRange(partitionKey, sortKeyLow, sortKeyHigh string) ([]map[string]*dynamodb.AttributeValue, error) {
	attributeValue := map[string]*dynamodb.AttributeValue{
		":partitionKeyVal": {
			S: aws.String(partitionKey),
		},
		":sortKeyLow": {
			S: aws.String(sortKeyLow),
		},
		":sortKeyHigh": {
			S: aws.String(sortKeyHigh),
		},
	}

	query := fmt.Sprintf("%s = :partitionKeyVal AND %s BETWEEN :sortKeyLow AND :sortKeyHigh", PartitionKeyFieldName, SortKeyFieldName)

	var queryInput = &dynamodb.QueryInput{
		Limit:                     &r.Limit,
		TableName:                 &r.table,
		KeyConditionExpression:    &query,
		ExpressionAttributeValues: attributeValue,
	}

	logrus.WithFields(logrus.Fields{
		"query":        query,
		"partitionKey": partitionKey,
		"sortKeyLow":   sortKeyLow,
		"sortKeyHigh":  sortKeyHigh,
		"limit":        r.Limit,
	}).Debug("query info")

	resp, err := r.db.Query(queryInput)
	itemsMapList := resp.Items

	moreRecords := true
	for moreRecords {
		if len(resp.LastEvaluatedKey) <= 0 {
			moreRecords = false
			break
		}

		queryInput.ExclusiveStartKey = resp.LastEvaluatedKey

		logrus.WithFields(logrus.Fields{
			"query":            query,
			"partitionKey":     partitionKey,
			"sortKeyLow":       sortKeyLow,
			"sortKeyHigh":      sortKeyHigh,
			"limit":            r.Limit,
			"lastEvaluatedKey": resp.LastEvaluatedKey,
		}).Debug("query info")

		resp, err = r.db.Query(queryInput)
		itemsMapList = append(itemsMapList, resp.Items...)

		if err != nil {
			return nil, err
		}
	}

	return itemsMapList, nil
}


type Location struct {
	RecordType    int
	RecordVersion int
	ApplicationID string
	ReportedTime  string
	Latitude      float64
	Longitude     float64
}

func (r *DynamoDB) ReadLocation(deviceID string, daysInPast int) ([]Location, error) {
	partitionKey := deviceID

	currentTimeStr := time.Now().UTC().Format("2006-01-02T15:04:05Z")

	past := time.Now().AddDate(0, 0, -daysInPast).UTC()
	pastUTC := time.Date(past.Year(), past.Month(), past.Day(), 0, 0, 0, 0, time.UTC)
	pastUTCStr := pastUTC.Format("2006-01-02T15:04:00Z")

	sortKeyLow := "01#" + pastUTCStr
	sortKeyHigh := "01#" + currentTimeStr

	locationRecords, err := r.queryBySortKeyRange(partitionKey, sortKeyLow, sortKeyHigh)
	if err != nil {
		return nil, fmt.Errorf("error reading location data from dynamo db=%s", err.Error())
	}

	var locations []Location
	for _, locationRecord := range locationRecords {
		location := Location{}
		err = dynamodbattribute.UnmarshalMap(locationRecord, &location)
		if err != nil {
			return nil, fmt.Errorf("error parsing dynamo db query records=%s", err.Error())
		}
		locations = append(locations, location)
	}
	return locations, nil
}
