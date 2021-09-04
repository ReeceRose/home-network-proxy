// Lambda function code

const { DynamoDBClient, ScanCommand } = require("@aws-sdk/client-dynamodb");

exports.handler = async (event) => {
  const userId = event.requestContext.authorizer.claims.sub;

  const { TABLE_NAME } = process.env;
  const headers = {
    "Content-Type": "application/json",
  };

  const client = new DynamoDBClient({ region: "us-east-1" });
  const command = new ScanCommand({
    TableName: TABLE_NAME,
    FilterExpression: "UserId = :userId",
    ExpressionAttributeValues: {
      ":userId": { S: userId },
    },
  });

  try {
    const results = await client.send(command);
    return {
      statusCode: 200,
      headers: headers,
      body: JSON.stringify({
        data: results.Items,
      }),
    };
  } catch (err) {
    console.error(err);
    return {
      statusCode: 500,
      headers: headers,
      body: JSON.stringify({
        error: err.message,
      }),
    };
  }
};
