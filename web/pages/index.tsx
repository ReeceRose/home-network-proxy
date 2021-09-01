import Amplify from 'aws-amplify';
// import Amplify, { Auth } from 'aws-amplify';
import awsconfig from '../aws-exports';
Amplify.configure(awsconfig);

import { withAuthenticator, AmplifySignOut } from '@aws-amplify/ui-react';

const Home = (): JSX.Element => {
  return (
    <div>
      <AmplifySignOut />
    </div>
  );
};

export default withAuthenticator(Home);
