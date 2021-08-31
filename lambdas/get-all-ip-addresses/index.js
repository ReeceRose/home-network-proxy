// Lambda function code

const { DynamoDBClient, ScanCommand } = require("@aws-sdk/client-dynamodb");

exports.handler = async (event) => {
  const userID = "123"; //TODO: read from auth

  const { TABLE_NAME } = process.env;

  const client = new DynamoDBClient({ region: "us-east-1" });
  const command = new ScanCommand({
    TableName: TABLE_NAME,
    FilterExpression: "UserId = :userId",
    ExpressionAttributeValues: {
      ":userId": { S: userID },
    },
  });

  try {
    const results = await client.send(command);
    return {
      statusCode: 200,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        data: results.Items,
      }),
    };
  } catch (err) {
    console.error(err);
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
