// Lambda function code

const { v4: uuidv4 } = require("uuid");

const {
  DynamoDBClient,
  GetItemCommand,
  UpdateItemCommand,
  PutItemCommand,
} = require("@aws-sdk/client-dynamodb");

exports.handler = async (event) => {
  const body = JSON.parse(event.body);
  // Not too much data validation going on, if any further data validation errors it will just return an error
  if (body === null || body.externalIP === null) {
    return {
      statusCode: 400,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        data: "invalid request body",
      }),
    };
  }

  const { TABLE_NAME } = process.env;

  const client = new DynamoDBClient({ region: "us-east-1" });

  try {
    if (body.id) {
      let command = new GetItemCommand({
        TableName: TABLE_NAME,
        Key: {
          Id: { S: body.id },
        },
      });
      let result = await client.send(command);
      if (result.Item) {
        // Note: technically UpdateItemCommand can be used for both insert/update but I think this way is easier to read
        command = new UpdateItemCommand({
          TableName: TABLE_NAME,
          Key: {
            Id: { S: body.id },
          },
          UpdateExpression: "set ExternalIP = :externalIP, Updated = :updated",
          ExpressionAttributeValues: {
            ":externalIP": { S: body.externalIP },
            ":updated": { S: Date.now().toString() },
          },
          ReturnValues: "ALL_NEW",
        });

        result = await client.send(command);
        return {
          statusCode: result.$metadata.httpStatusCode,
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            data: result.Attributes,
          }),
        };
      }
    }

    const now = Date.now().toString();
    command = new PutItemCommand({
      TableName: TABLE_NAME,
      Item: {
        Id: { S: uuidv4() },
        Created: { S: now },
        Updated: { S: now },
        ExternalIP: { S: body.externalIP },
        UserId: { S: body.userId },
      },
    });

    result = await client.send(command);
    return {
      statusCode: result.$metadata.httpStatusCode,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        success: true,
      }),
    };
  } catch (err) {
    return {
      statusCode: 500,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        error: err.message,
      }),
    };
  }
};
