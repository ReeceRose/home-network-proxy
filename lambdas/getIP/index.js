// Lambda function code

const { DynamoDBClient, QueryCommand } = require("@aws-sdk/client-dynamodb");

exports.handler = async (event) => {
  if (
    !event.queryStringParameters ||
    !event.queryStringParameters["externalIP"]
  ) {
    return {
      statusCode: 400,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        message: "please specify externalIP in query string",
      }),
    };
  }
  externalIP = event.queryStringParameters["externalIP"];

  const { TABLE_NAME } = process.env;

  const client = new DynamoDBClient({ region: "us-east-1" });
  const command = new QueryCommand({
    TableName: TABLE_NAME,
    KeyConditionExpression: "#ExternalIP = :externalIP",
    ExpressionAttributeValues: {
      ":externalIP": externalIP,
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
        message: results.Items,
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
        message: err.message,
      }),
    };
  }
};
