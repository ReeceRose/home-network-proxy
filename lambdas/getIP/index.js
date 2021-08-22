// Lambda function code

const { DynamoDBClient, QueryCommand } = require("@aws-sdk/client-dynamodb");

exports.handler = async (event) => {
  if (
    !event.queryStringParameters ||
    !event.queryStringParameters["id"] ||
    !event.queryStringParameters["userID"]
  ) {
    return {
      statusCode: 400,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        error: "please specify id and userID in query string",
      }),
    };
  }
  Id = event.queryStringParameters["id"];
  userID = event.queryStringParameters["userID"];

  const { TABLE_NAME } = process.env;

  const client = new DynamoDBClient({ region: "us-east-1" });
  const command = new QueryCommand({
    TableName: TABLE_NAME,
    KeyConditionExpression: "Id = :Id",
    ExpressionAttributeValues: {
      ":Id": { S: Id },
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
