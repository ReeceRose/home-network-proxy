import { useEffect, useState } from 'react';

import Amplify, { Auth } from 'aws-amplify';
import awsconfig from '../aws-exports';
Amplify.configure(awsconfig);

import { withAuthenticator, AmplifySignOut } from '@aws-amplify/ui-react';
import axios from 'axios';
import { IPAddress } from '../interfaces/Index';

import Navbar from '../components/Navbar';
import Footer from '../components/Footer';
import ReportedIPAddress from '../components/Tables/ReportedIPAddress';

const Home = (): JSX.Element => {
  const [addresses, setAddresses] = useState<IPAddress[]>();

  useEffect(() => {
    async function fetchIPAddresses() {
      const user = await Auth.currentAuthenticatedUser();
      const token = user.signInUserSession.idToken.jwtToken;
      const response = await axios.get(
        'https://p7xsh0rld4.execute-api.us-east-1.amazonaws.com/production/ip/all',
        {
          headers: {
            Authorization: 'Bearer ' + token,
          },
        }
      );

      const addresses: IPAddress[] = JSON.parse(
        JSON.stringify(response.data.data)
      );
      setAddresses(addresses);
    }

    fetchIPAddresses();
  }, []);

  const deleteIP = (id: string) => {
    async function deleteIPAddress() {
      const user = await Auth.currentAuthenticatedUser();
      const token = user.signInUserSession.idToken.jwtToken;
      const response = await axios.delete(
        'https://p7xsh0rld4.execute-api.us-east-1.amazonaws.com/production/ip',
        {
          headers: {
            Authorization: 'Bearer ' + token,
          },
          data: JSON.stringify({ id: id }),
        }
      );
      if (response.status !== 200) {
        location.reload();
      }
      setAddresses(addresses?.filter((address) => address.Id.S !== id));
    }
    deleteIPAddress();
  };

  return (
    <div className="relative flex flex-col w-full min-h-screen bg-gray-100">
      <div>
        <AmplifySignOut />
        <Navbar />
        <div className="flex-grow w-full px-4 m-24 mx-auto mb-0 md:px-10">
          <div className="flex flex-wrap">
            <div className="w-full px-4 mb-12 xl:mb-0">
              {addresses?.map((address) => (
                <ReportedIPAddress
                  key={address.Id.S}
                  address={address}
                  deleteIP={() => deleteIP(address.Id.S)}
                />
              ))}
            </div>
          </div>
        </div>
        <Footer />
      </div>
    </div>
  );
};

export default withAuthenticator(Home);
