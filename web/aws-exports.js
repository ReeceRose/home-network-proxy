const awsmobile = {
  aws_project_region: process.env.project_region,
  aws_cognito_identity_pool_id: process.env.cognito_identity_pool_id,
  aws_cognito_region: process.env.cognito_region,
  aws_user_pools_id: process.env._user_pools_id,
  aws_user_pools_web_client_id: process.env.user_pools_web_client_id,
  oauth: {},
};

export default awsmobile;
