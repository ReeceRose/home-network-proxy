// Lambda function code

const { DynamoDBClient, QueryCommand } = require("@aws-sdk/client-dynamodb");

exports.handler = async (event) => {
  if (!event.queryStringParameters || !event.queryStringParameters["userID"]) {
    return {
      statusCode: 400,
      headers: {
        "Content-Type": "application/json",
      },
      data: JSON.stringify({
        error: "please specify userID in query string",
      }),
    };
  }
  userID = event.queryStringParameters["userID"];

  const { TABLE_NAME } = process.env;

  const client = new DynamoDBClient({ region: "us-east-1" });
  const command = new QueryCommand({
    TableName: TABLE_NAME,
    KeyConditionExpression: "#UserID = :userID",
    ExpressionAttributeValues: {
      ":userID": userID,
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
