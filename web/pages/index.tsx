import { useEffect, useState } from 'react';

import Amplify, { Auth } from 'aws-amplify';
import awsconfig from '../aws-exports';
Amplify.configure(awsconfig);

import { withAuthenticator, AmplifySignOut } from '@aws-amplify/ui-react';
import axios from 'axios';
import { IPAddress } from '../interfaces/Index';

const Home = (): JSX.Element => {
  const [addresses, setAddresses] = useState<IPAddress[]>();
  useEffect(() => {
    async function fetchIPAddresses() {
      const user = await Auth.currentAuthenticatedUser();
      console.log(user);
      const token = user.signInUserSession.idToken.jwtToken;
      const response = await axios.get(
        'https://p7xsh0rld4.execute-api.us-east-1.amazonaws.com/production/ip/all',
        {
          headers: {
            Authorization: 'Bearer ' + token,
          },
        }
      );
      console.log(response);
      const addresses: IPAddress[] = JSON.parse(
        JSON.stringify(response.data.data)
      );
      setAddresses(addresses);
    }

    fetchIPAddresses();
  }, []);

  return (
    <div>
      <AmplifySignOut />
      {addresses?.map((address) => (
        <div key={address.Id.S}>
          <p>
            {address.ExternalIP.S} |
            {new Date(parseInt(address.Updated.S)).toString()}
          </p>
        </div>
      ))}
    </div>
  );
};

export default withAuthenticator(Home);
