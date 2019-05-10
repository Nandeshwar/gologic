func (r *DynamoDB) PutArbitraryJSON(arbitraryJSON map[string]interface{}) (*dynamodb.PutItemOutput, error) {

	err := r.verifyRequiredAttributes(arbitraryJSON)
	if err != nil {
		return nil, err
	}

	dynamoObj, err := dynamodbattribute.MarshalMap(arbitraryJSON)
	if err != nil {
		return nil, err
	}

	var tableName = r.table
	var conditionExpression = "attribute_not_exists(pkeyDeviceId)" //Prevents accidental overwrites
	var returnValues = "NONE"
	var returnConsumedCapacity = "TOTAL"

	input := &dynamodb.PutItemInput{
		TableName:              &tableName,
		ConditionExpression:    &conditionExpression,
		Item:                   dynamoObj,
		ReturnValues:           &returnValues,
		ReturnConsumedCapacity: &returnConsumedCapacity,
	}

	return r.db.PutItem(input)
}


func (r *DynamoDB) verifyRequiredAttributes(arbitraryJSON map[string]interface{}) error {
	//Note: Only a few attributes are explicitly verified at time of insertion.
	//The full schema IS NOT verified at time of insertion. It's assumed upstream producer has done schema verification

	partitionKey := arbitraryJSON[PartitionKeyFieldName]
	if partitionKey == "" {
		return errors.New(fmt.Sprintf("Malformed dynamodb item. Missing partition key %s", PartitionKeyFieldName))
	}
	switch partitionKey.(type) {
	case string:
	default:
		return errors.New("Malformed dynamodb item. Partition key was not a string.")
	}

	sortKey := arbitraryJSON[SortKeyFieldName]
	if sortKey == "" {
		return errors.New(fmt.Sprintf("Malformed dynamodb item. Missing sort key %s.", SortKeyFieldName))
	}
	switch sortKey.(type) {
	case string:
	default:
		return errors.New("Malformed dynamodb item. Sort key was not a string.")
	}

	ttl := arbitraryJSON[TTLFieldName]
	if ttl == "" {
		return errors.New(fmt.Sprintf("Malformed dynamodb item. Missing ttl attribute %s.", TTLFieldName))
	}

	switch ttl.(type) {
	case time.Time:
	default:
		return errors.New("Malformed dynamodb item. TTL attribute was not a time.")
	}

	logrus.WithFields(logrus.Fields{
		"partitionKey":  partitionKey,
		"sortKey":       sortKey,
		"expirationKey": ttl,
	}).Debug("Verified required attributes exist for item.")

	return nil
}
