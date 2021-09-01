import { useState } from 'react';

import Amplify, { Auth } from 'aws-amplify';
import awsconfig from '../aws-exports';
Amplify.configure(awsconfig);

import { withAuthenticator, AmplifySignOut } from '@aws-amplify/ui-react';

const Home = (): JSX.Element => {
  const [userId, setUserId] = useState();
  Auth.currentUserInfo().then((data) => {
    setUserId(data.attributes.sub);
  });
  return (
    <div>
      <AmplifySignOut />
      {userId && <h1>{userId}</h1>}
    </div>
  );
};

export default withAuthenticator(Home);
