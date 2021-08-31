// Lambda function code

exports.handler = async (event) => {
  return {
    statusCode: 200,
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      data: "ok",
    }),
  };
};
